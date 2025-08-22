// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnRealTimeLogProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnRealTimeLogProjectResponseBody
	GetRequestId() *string
}

type DeleteDcdnRealTimeLogProjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnRealTimeLogProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnRealTimeLogProjectResponseBody) SetRequestId(v string) *DeleteDcdnRealTimeLogProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
