// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBdrcServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBdrcServiceResponseBodyData) *GetBdrcServiceResponseBody
	GetData() *GetBdrcServiceResponseBodyData
	SetRequestId(v string) *GetBdrcServiceResponseBody
	GetRequestId() *string
}

type GetBdrcServiceResponseBody struct {
	// The data that is returned if the call is successful.
	Data *GetBdrcServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identity of the request.
	//
	// example:
	//
	// 5748C531-80B1-5C31-8421-63A1830B9E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBdrcServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponseBody) GetData() *GetBdrcServiceResponseBodyData {
	return s.Data
}

func (s *GetBdrcServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBdrcServiceResponseBody) SetData(v *GetBdrcServiceResponseBodyData) *GetBdrcServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetBdrcServiceResponseBody) SetRequestId(v string) *GetBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBdrcServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBdrcServiceResponseBodyData struct {
	// The time when the service was enabled (UNIX timestamp).
	//
	// example:
	//
	// 1726169608
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// The time when the data protection score was updated (UNIX timestamp).
	//
	// example:
	//
	// 1726169608
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// The initialization status of the service.
	//
	// example:
	//
	// SUCCESS
	ServiceInitializeStatus *string `json:"ServiceInitializeStatus,omitempty" xml:"ServiceInitializeStatus,omitempty"`
	// The enabling status of the service.
	//
	// example:
	//
	// OPENED
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetBdrcServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBdrcServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponseBodyData) GetOpenTime() *int64 {
	return s.OpenTime
}

func (s *GetBdrcServiceResponseBodyData) GetProtectionScoreUpdatedTime() *int64 {
	return s.ProtectionScoreUpdatedTime
}

func (s *GetBdrcServiceResponseBodyData) GetServiceInitializeStatus() *string {
	return s.ServiceInitializeStatus
}

func (s *GetBdrcServiceResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetBdrcServiceResponseBodyData) SetOpenTime(v int64) *GetBdrcServiceResponseBodyData {
	s.OpenTime = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetProtectionScoreUpdatedTime(v int64) *GetBdrcServiceResponseBodyData {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetServiceInitializeStatus(v string) *GetBdrcServiceResponseBodyData {
	s.ServiceInitializeStatus = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetServiceStatus(v string) *GetBdrcServiceResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
