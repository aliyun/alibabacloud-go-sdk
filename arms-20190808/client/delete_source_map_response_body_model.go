// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteSourceMapResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteSourceMapResponseBody
	GetRequestId() *string
}

type DeleteSourceMapResponseBody struct {
	// Indicates whether the SourceMap files are deleted. Valid values:
	//
	// 	- success: The SourceMap files are deleted.
	//
	// 	- false: The SourceMap files fail to be deleted.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSourceMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteSourceMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceMapResponseBody) SetData(v string) *DeleteSourceMapResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSourceMapResponseBody) SetRequestId(v string) *DeleteSourceMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceMapResponseBody) Validate() error {
	return dara.Validate(s)
}
