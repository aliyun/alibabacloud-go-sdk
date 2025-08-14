// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *RemoveEventResponseBody
	GetResultObject() *bool
}

type RemoveEventResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s RemoveEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEventResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *RemoveEventResponseBody) SetRequestId(v string) *RemoveEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEventResponseBody) SetResultObject(v bool) *RemoveEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RemoveEventResponseBody) Validate() error {
	return dara.Validate(s)
}
