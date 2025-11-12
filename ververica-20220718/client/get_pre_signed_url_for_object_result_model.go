// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreSignedUrlForObjectResult interface {
	dara.Model
	String() string
	GoString() string
	SetJarUrl(v string) *GetPreSignedUrlForObjectResult
	GetJarUrl() *string
	SetPreSignedUrl(v string) *GetPreSignedUrlForObjectResult
	GetPreSignedUrl() *string
}

type GetPreSignedUrlForObjectResult struct {
	JarUrl       *string `json:"jarUrl,omitempty" xml:"jarUrl,omitempty"`
	PreSignedUrl *string `json:"preSignedUrl,omitempty" xml:"preSignedUrl,omitempty"`
}

func (s GetPreSignedUrlForObjectResult) String() string {
	return dara.Prettify(s)
}

func (s GetPreSignedUrlForObjectResult) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlForObjectResult) GetJarUrl() *string {
	return s.JarUrl
}

func (s *GetPreSignedUrlForObjectResult) GetPreSignedUrl() *string {
	return s.PreSignedUrl
}

func (s *GetPreSignedUrlForObjectResult) SetJarUrl(v string) *GetPreSignedUrlForObjectResult {
	s.JarUrl = &v
	return s
}

func (s *GetPreSignedUrlForObjectResult) SetPreSignedUrl(v string) *GetPreSignedUrlForObjectResult {
	s.PreSignedUrl = &v
	return s
}

func (s *GetPreSignedUrlForObjectResult) Validate() error {
	return dara.Validate(s)
}
