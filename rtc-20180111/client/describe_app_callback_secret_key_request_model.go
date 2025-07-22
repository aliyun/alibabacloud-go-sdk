// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallbackSecretKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppCallbackSecretKeyRequest
	GetAppId() *string
}

type DescribeAppCallbackSecretKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppCallbackSecretKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallbackSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppCallbackSecretKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppCallbackSecretKeyRequest) SetAppId(v string) *DescribeAppCallbackSecretKeyRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppCallbackSecretKeyRequest) Validate() error {
	return dara.Validate(s)
}
