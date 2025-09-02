// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CheckMetaTableResponseBody
	GetData() *bool
	SetRequestId(v string) *CheckMetaTableResponseBody
	GetRequestId() *string
}

type CheckMetaTableResponseBody struct {
	// Indicates whether the metatable exists.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMetaTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaTableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMetaTableResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckMetaTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMetaTableResponseBody) SetData(v bool) *CheckMetaTableResponseBody {
	s.Data = &v
	return s
}

func (s *CheckMetaTableResponseBody) SetRequestId(v string) *CheckMetaTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMetaTableResponseBody) Validate() error {
	return dara.Validate(s)
}
