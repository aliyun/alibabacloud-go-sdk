// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenConsoleSlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenConsoleSlsResponseBody
	GetRequestId() *string
	SetResultObject(v string) *OpenConsoleSlsResponseBody
	GetResultObject() *string
}

type OpenConsoleSlsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *string `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s OpenConsoleSlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenConsoleSlsResponseBody) GoString() string {
	return s.String()
}

func (s *OpenConsoleSlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenConsoleSlsResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *OpenConsoleSlsResponseBody) SetRequestId(v string) *OpenConsoleSlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenConsoleSlsResponseBody) SetResultObject(v string) *OpenConsoleSlsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *OpenConsoleSlsResponseBody) Validate() error {
	return dara.Validate(s)
}
