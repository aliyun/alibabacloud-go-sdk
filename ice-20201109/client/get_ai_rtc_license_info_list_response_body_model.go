// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcLicenseInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiRtcLicenseInfoListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAiRtcLicenseInfoListResponseBody
	GetHttpStatusCode() *int32
	SetLicenseInfoList(v []*AiRtcLicenseInfoDTO) *GetAiRtcLicenseInfoListResponseBody
	GetLicenseInfoList() []*AiRtcLicenseInfoDTO
	SetMessage(v string) *GetAiRtcLicenseInfoListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiRtcLicenseInfoListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiRtcLicenseInfoListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetAiRtcLicenseInfoListResponseBody
	GetTotalCount() *int64
}

type GetAiRtcLicenseInfoListResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// An array of AiRtcLicenseInfoDTO objects, each representing a license batch.
	LicenseInfoList []*AiRtcLicenseInfoDTO `json:"LicenseInfoList,omitempty" xml:"LicenseInfoList,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAiRtcLicenseInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcLicenseInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetLicenseInfoList() []*AiRtcLicenseInfoDTO {
	return s.LicenseInfoList
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiRtcLicenseInfoListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetCode(v string) *GetAiRtcLicenseInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetHttpStatusCode(v int32) *GetAiRtcLicenseInfoListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetLicenseInfoList(v []*AiRtcLicenseInfoDTO) *GetAiRtcLicenseInfoListResponseBody {
	s.LicenseInfoList = v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetMessage(v string) *GetAiRtcLicenseInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetRequestId(v string) *GetAiRtcLicenseInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetSuccess(v bool) *GetAiRtcLicenseInfoListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) SetTotalCount(v int64) *GetAiRtcLicenseInfoListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponseBody) Validate() error {
	if s.LicenseInfoList != nil {
		for _, item := range s.LicenseInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
