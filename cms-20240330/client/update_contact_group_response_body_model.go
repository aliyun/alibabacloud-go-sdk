// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateContactGroupResponseBody
	GetRequestId() *string
}

type UpdateContactGroupResponseBody struct {
	// The unique ID of the request, which is used for log tracing and troubleshooting.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactGroupResponseBody) SetRequestId(v string) *UpdateContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
