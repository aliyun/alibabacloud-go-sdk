// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRepoTagScanStatusResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetRepoTagScanStatusResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetRepoTagScanStatusResponseBody
	GetRequestId() *string
	SetScanService(v string) *GetRepoTagScanStatusResponseBody
	GetScanService() *string
	SetStatus(v string) *GetRepoTagScanStatusResponseBody
	GetStatus() *string
}

type GetRepoTagScanStatusResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the scanning engine.
	//
	// 	- `ACR_SCAN_SERVICE`: Trivy scan engine provided by Container Registry
	//
	// 	- `SAS_SCAN_SERVICE`: Security Center scan engine
	//
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// The scanning status of the image tag. Valid values:
	//
	// 	- `SCANNING`: The image tag is being scanned.
	//
	// 	- `COMPLETE`: The scanning of the image tag is complete.
	//
	// 	- `FAILED`: The image tag failed to be scanned.
	//
	// 	- `RETRYING`: The system is retrying to scan the image tag.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRepoTagScanStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoTagScanStatusResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoTagScanStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoTagScanStatusResponseBody) GetScanService() *string {
	return s.ScanService
}

func (s *GetRepoTagScanStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRepoTagScanStatusResponseBody) SetCode(v string) *GetRepoTagScanStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetIsSuccess(v bool) *GetRepoTagScanStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetRequestId(v string) *GetRepoTagScanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetScanService(v string) *GetRepoTagScanStatusResponseBody {
	s.ScanService = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetStatus(v string) *GetRepoTagScanStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
