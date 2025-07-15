// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessAndAbolishApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *RemoveVpcAccessAndAbolishApisResponseBody
	GetOperationId() *string
	SetRequestId(v string) *RemoveVpcAccessAndAbolishApisResponseBody
	GetRequestId() *string
}

type RemoveVpcAccessAndAbolishApisResponseBody struct {
	// example:
	//
	// f7834d74be4e41aa8e607b0fafae9b33
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVpcAccessAndAbolishApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) SetOperationId(v string) *RemoveVpcAccessAndAbolishApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) SetRequestId(v string) *RemoveVpcAccessAndAbolishApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) Validate() error {
	return dara.Validate(s)
}
