resource "aws_s3_bucket" "bucket" {
  bucket = "bucket-for-terratest"
  tags = {
    Name = "Terratest"
  }
}
