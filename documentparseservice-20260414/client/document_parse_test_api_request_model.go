// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseTestApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *DocumentParseTestApiRequest
	GetImageUrl() *string
	SetType(v int64) *DocumentParseTestApiRequest
	GetType() *int64
}

type DocumentParseTestApiRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1x2xCAxz1gK0jSZSgXXavwpXa-3541-662.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// MAIN_FLOW
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DocumentParseTestApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseTestApiRequest) GoString() string {
	return s.String()
}

func (s *DocumentParseTestApiRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DocumentParseTestApiRequest) GetType() *int64 {
	return s.Type
}

func (s *DocumentParseTestApiRequest) SetImageUrl(v string) *DocumentParseTestApiRequest {
	s.ImageUrl = &v
	return s
}

func (s *DocumentParseTestApiRequest) SetType(v int64) *DocumentParseTestApiRequest {
	s.Type = &v
	return s
}

func (s *DocumentParseTestApiRequest) Validate() error {
	return dara.Validate(s)
}
