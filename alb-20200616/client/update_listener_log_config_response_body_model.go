// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateListenerLogConfigResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateListenerLogConfigResponseBody
	GetRequestId() *string
}

type UpdateListenerLogConfigResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateListenerLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateListenerLogConfigResponseBody) SetJobId(v string) *UpdateListenerLogConfigResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerLogConfigResponseBody) SetRequestId(v string) *UpdateListenerLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateListenerLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
