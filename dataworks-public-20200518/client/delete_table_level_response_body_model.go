// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteResult(v bool) *DeleteTableLevelResponseBody
	GetDeleteResult() *bool
	SetRequestId(v string) *DeleteTableLevelResponseBody
	GetRequestId() *string
}

type DeleteTableLevelResponseBody struct {
	// Indicates whether the table level was deleted.
	//
	// example:
	//
	// true
	DeleteResult *bool `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTableLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableLevelResponseBody) GetDeleteResult() *bool {
	return s.DeleteResult
}

func (s *DeleteTableLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableLevelResponseBody) SetDeleteResult(v bool) *DeleteTableLevelResponseBody {
	s.DeleteResult = &v
	return s
}

func (s *DeleteTableLevelResponseBody) SetRequestId(v string) *DeleteTableLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
