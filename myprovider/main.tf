provider "github" {
  token = "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"
}

resource "github_repository" "example" {
  name        = "swatiterratest"
  description = "mgc 88 is a gambling website backend developed by walking dreamz"
  visibility  = "public"
}