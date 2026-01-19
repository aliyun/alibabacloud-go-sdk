// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *CreateCloudAccountResponseBody
	GetCloudAccountId() *string
	SetRequestId(v string) *CreateCloudAccountResponseBody
	GetRequestId() *string
}

type CreateCloudAccountResponseBody struct {
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseBody) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *CreateCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudAccountResponseBody) SetCloudAccountId(v string) *CreateCloudAccountResponseBody {
	s.CloudAccountId = &v
	return s
}

func (s *CreateCloudAccountResponseBody) SetRequestId(v string) *CreateCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
