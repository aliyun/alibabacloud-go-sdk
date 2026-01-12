// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *QueryProjectDetailRequest
	GetJwtToken() *string
	SetOssKey(v string) *QueryProjectDetailRequest
	GetOssKey() *string
	SetProjectId(v string) *QueryProjectDetailRequest
	GetProjectId() *string
}

type QueryProjectDetailRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	OssKey    *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *QueryProjectDetailRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *QueryProjectDetailRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectDetailRequest) SetJwtToken(v string) *QueryProjectDetailRequest {
	s.JwtToken = &v
	return s
}

func (s *QueryProjectDetailRequest) SetOssKey(v string) *QueryProjectDetailRequest {
	s.OssKey = &v
	return s
}

func (s *QueryProjectDetailRequest) SetProjectId(v string) *QueryProjectDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectDetailRequest) Validate() error {
	return dara.Validate(s)
}
