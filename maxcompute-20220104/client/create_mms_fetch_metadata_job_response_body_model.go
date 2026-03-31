// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsFetchMetadataJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMmsFetchMetadataJobResponseBodyData) *CreateMmsFetchMetadataJobResponseBody
	GetData() *CreateMmsFetchMetadataJobResponseBodyData
	SetRequestId(v string) *CreateMmsFetchMetadataJobResponseBody
	GetRequestId() *string
}

type CreateMmsFetchMetadataJobResponseBody struct {
	Data *CreateMmsFetchMetadataJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// CC4D05E8-0613-5A8E-9339-A0EBD097A69E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponseBody) GetData() *CreateMmsFetchMetadataJobResponseBodyData {
	return s.Data
}

func (s *CreateMmsFetchMetadataJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMmsFetchMetadataJobResponseBody) SetData(v *CreateMmsFetchMetadataJobResponseBodyData) *CreateMmsFetchMetadataJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsFetchMetadataJobResponseBody) SetRequestId(v string) *CreateMmsFetchMetadataJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMmsFetchMetadataJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMmsFetchMetadataJobResponseBodyData struct {
	// example:
	//
	// 1000002
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponseBodyData) GetScanId() *int64 {
	return s.ScanId
}

func (s *CreateMmsFetchMetadataJobResponseBodyData) SetScanId(v int64) *CreateMmsFetchMetadataJobResponseBodyData {
	s.ScanId = &v
	return s
}

func (s *CreateMmsFetchMetadataJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
