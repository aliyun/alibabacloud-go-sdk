// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQueryResultResponseBody
	GetRequestId() *string
}

type ListQueryResultResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 9E5BCFAA-98B3-51D0-9188-B1BC07589337
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQueryResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueryResultResponseBody) SetRequestId(v string) *ListQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryResultResponseBody) Validate() error {
	return dara.Validate(s)
}
