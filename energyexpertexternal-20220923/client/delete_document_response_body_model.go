// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteDocumentResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteDocumentResponseBody
	GetRequestId() *string
}

type DeleteDocumentResponseBody struct {
	// Returns true on success, false otherwise
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentResponseBody) SetData(v bool) *DeleteDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
