// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteResourceInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceInstancesResponseBody
	GetRequestId() *string
}

type DeleteResourceInstancesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Instances eas-i-011227132046,eas-i-011227132046 are deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceInstancesResponseBody) SetMessage(v string) *DeleteResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceInstancesResponseBody) SetRequestId(v string) *DeleteResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
