// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DocumentDetailItem) *DetailDocumentResponseBody
	GetData() *DocumentDetailItem
	SetRequestId(v string) *DetailDocumentResponseBody
	GetRequestId() *string
}

type DetailDocumentResponseBody struct {
	// example:
	//
	// true
	Data *DocumentDetailItem `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetailDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DetailDocumentResponseBody) GetData() *DocumentDetailItem {
	return s.Data
}

func (s *DetailDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailDocumentResponseBody) SetData(v *DocumentDetailItem) *DetailDocumentResponseBody {
	s.Data = v
	return s
}

func (s *DetailDocumentResponseBody) SetRequestId(v string) *DetailDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
