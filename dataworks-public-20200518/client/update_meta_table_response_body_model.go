// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaTableResponseBody
	GetRequestId() *string
	SetUpdateResult(v bool) *UpdateMetaTableResponseBody
	GetUpdateResult() *bool
}

type UpdateMetaTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the metadata information about the table was updated.
	//
	// example:
	//
	// true
	UpdateResult *bool `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty"`
}

func (s UpdateMetaTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaTableResponseBody) GetUpdateResult() *bool {
	return s.UpdateResult
}

func (s *UpdateMetaTableResponseBody) SetRequestId(v string) *UpdateMetaTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaTableResponseBody) SetUpdateResult(v bool) *UpdateMetaTableResponseBody {
	s.UpdateResult = &v
	return s
}

func (s *UpdateMetaTableResponseBody) Validate() error {
	return dara.Validate(s)
}
