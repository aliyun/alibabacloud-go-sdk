// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *DescribeAuthVerifyRequest
	GetCertifyId() *string
	SetSceneId(v int64) *DescribeAuthVerifyRequest
	GetSceneId() *int64
}

type DescribeAuthVerifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// shsf57a4e0d9981c3bd66dc754f3d3cd
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100000****
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeAuthVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeAuthVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeAuthVerifyRequest) SetCertifyId(v string) *DescribeAuthVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeAuthVerifyRequest) SetSceneId(v int64) *DescribeAuthVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeAuthVerifyRequest) Validate() error {
	return dara.Validate(s)
}
