// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetProjectDatasetRequest
	GetJwtToken() *string
	SetOssKey(v string) *GetProjectDatasetRequest
	GetOssKey() *string
}

type GetProjectDatasetRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s GetProjectDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetProjectDatasetRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetProjectDatasetRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *GetProjectDatasetRequest) SetJwtToken(v string) *GetProjectDatasetRequest {
	s.JwtToken = &v
	return s
}

func (s *GetProjectDatasetRequest) SetOssKey(v string) *GetProjectDatasetRequest {
	s.OssKey = &v
	return s
}

func (s *GetProjectDatasetRequest) Validate() error {
	return dara.Validate(s)
}
