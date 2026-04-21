// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListApiInfosResponseBody
	GetRequestId() *string
}

type ListApiInfosResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 86CD53FA-81A5-56A0-AE6F-CA1C56F48574
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApiInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiInfosResponseBody) SetRequestId(v string) *ListApiInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiInfosResponseBody) Validate() error {
	return dara.Validate(s)
}
