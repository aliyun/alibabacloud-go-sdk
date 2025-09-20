// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareImageFacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompareImageFacesResponseBody
	GetRequestId() *string
	SetSimilarity(v float32) *CompareImageFacesResponseBody
	GetSimilarity() *float32
}

type CompareImageFacesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F73AC982-2B9E-4ECD-AED5-F8331C5******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The face similarity. A larger value indicates a higher face similarity. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.8848152756690983
	Similarity *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s CompareImageFacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareImageFacesResponseBody) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *CompareImageFacesResponseBody) SetRequestId(v string) *CompareImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSimilarity(v float32) *CompareImageFacesResponseBody {
	s.Similarity = &v
	return s
}

func (s *CompareImageFacesResponseBody) Validate() error {
	return dara.Validate(s)
}
