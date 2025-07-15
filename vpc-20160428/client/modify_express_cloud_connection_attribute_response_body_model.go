// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressCloudConnectionAttributeResponseBody
	GetRequestId() *string
}

type ModifyExpressCloudConnectionAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E6385514-B0CC-48E3-B9F9-F7BFF64460A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExpressCloudConnectionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressCloudConnectionAttributeResponseBody) SetRequestId(v string) *ModifyExpressCloudConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
