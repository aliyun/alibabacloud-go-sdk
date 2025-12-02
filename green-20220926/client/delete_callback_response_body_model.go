// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteCallbackResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteCallbackResponseBody
	GetRequestId() *string
}

type DeleteCallbackResponseBody struct {
	// Returned data.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCallbackResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCallbackResponseBody) SetData(v bool) *DeleteCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCallbackResponseBody) SetRequestId(v string) *DeleteCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
