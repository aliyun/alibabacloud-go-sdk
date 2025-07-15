// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSpMetadataResponseBody
	GetRequestId() *string
	SetSpMetadata(v string) *GetSpMetadataResponseBody
	GetSpMetadata() *string
}

type GetSpMetadataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the SP.
	SpMetadata *string `json:"SpMetadata,omitempty" xml:"SpMetadata,omitempty"`
}

func (s GetSpMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpMetadataResponseBody) GetSpMetadata() *string {
	return s.SpMetadata
}

func (s *GetSpMetadataResponseBody) SetRequestId(v string) *GetSpMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpMetadataResponseBody) SetSpMetadata(v string) *GetSpMetadataResponseBody {
	s.SpMetadata = &v
	return s
}

func (s *GetSpMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
