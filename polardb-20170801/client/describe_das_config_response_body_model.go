// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDasConfigResponseBody
	GetRequestId() *string
	SetStorageAutoScale(v string) *DescribeDasConfigResponseBody
	GetStorageAutoScale() *string
	SetStorageUpperBound(v int64) *DescribeDasConfigResponseBody
	GetStorageUpperBound() *int64
}

type DescribeDasConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593AE1C5-B70C-463F-9207-074639******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether to enable automatic storage scaling for the Standard Edition cluster. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The maximum storage capacity that is allowed for storage automatic scaling of the Standard Edition cluster. Unit: GB.
	//
	// >  This parameter is valid only when the StorageAutoScale parameter is set to Enable.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
}

func (s DescribeDasConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDasConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDasConfigResponseBody) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *DescribeDasConfigResponseBody) GetStorageUpperBound() *int64 {
	return s.StorageUpperBound
}

func (s *DescribeDasConfigResponseBody) SetRequestId(v string) *DescribeDasConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDasConfigResponseBody) SetStorageAutoScale(v string) *DescribeDasConfigResponseBody {
	s.StorageAutoScale = &v
	return s
}

func (s *DescribeDasConfigResponseBody) SetStorageUpperBound(v int64) *DescribeDasConfigResponseBody {
	s.StorageUpperBound = &v
	return s
}

func (s *DescribeDasConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
