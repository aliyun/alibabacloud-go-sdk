// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskHitDataUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetScannerTaskHitDataUrlResponseBodyData) *GetScannerTaskHitDataUrlResponseBody
	GetData() *GetScannerTaskHitDataUrlResponseBodyData
	SetRequestId(v string) *GetScannerTaskHitDataUrlResponseBody
	GetRequestId() *string
}

type GetScannerTaskHitDataUrlResponseBody struct {
	// The query result, which contains the task identifier and the temporary download URL for the hit data.
	Data *GetScannerTaskHitDataUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use this ID for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetScannerTaskHitDataUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskHitDataUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetScannerTaskHitDataUrlResponseBody) GetData() *GetScannerTaskHitDataUrlResponseBodyData {
	return s.Data
}

func (s *GetScannerTaskHitDataUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScannerTaskHitDataUrlResponseBody) SetData(v *GetScannerTaskHitDataUrlResponseBodyData) *GetScannerTaskHitDataUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetScannerTaskHitDataUrlResponseBody) SetRequestId(v string) *GetScannerTaskHitDataUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScannerTaskHitDataUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScannerTaskHitDataUrlResponseBodyData struct {
	// The temporary download URL for the attack hit data (hits.csv). The URL is valid for 2 hours. After the URL expires, call this operation again to obtain a new URL. If the task is an existing task or the agent execution mode does not produce hit data, this parameter is an empty string. The actual value is a signed temporary URL of Object Storage Service (OSS) that contains signature parameters.
	//
	// example:
	//
	// https://example.com/result.task-abc123def4567.hits.csv
	ScannerTaskHitDataDownloadUrl *string `json:"ScannerTaskHitDataDownloadUrl,omitempty" xml:"ScannerTaskHitDataDownloadUrl,omitempty"`
	// The scan task identifier echoed from the request.
	//
	// example:
	//
	// task-abc123def4567
	ScannerTaskId *string `json:"ScannerTaskId,omitempty" xml:"ScannerTaskId,omitempty"`
}

func (s GetScannerTaskHitDataUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskHitDataUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScannerTaskHitDataUrlResponseBodyData) GetScannerTaskHitDataDownloadUrl() *string {
	return s.ScannerTaskHitDataDownloadUrl
}

func (s *GetScannerTaskHitDataUrlResponseBodyData) GetScannerTaskId() *string {
	return s.ScannerTaskId
}

func (s *GetScannerTaskHitDataUrlResponseBodyData) SetScannerTaskHitDataDownloadUrl(v string) *GetScannerTaskHitDataUrlResponseBodyData {
	s.ScannerTaskHitDataDownloadUrl = &v
	return s
}

func (s *GetScannerTaskHitDataUrlResponseBodyData) SetScannerTaskId(v string) *GetScannerTaskHitDataUrlResponseBodyData {
	s.ScannerTaskId = &v
	return s
}

func (s *GetScannerTaskHitDataUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
