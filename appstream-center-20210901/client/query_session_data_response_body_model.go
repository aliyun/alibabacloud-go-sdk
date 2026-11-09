// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySessionDataResponseBody
	GetCode() *string
	SetIndexData(v []*QuerySessionDataResponseBodyIndexData) *QuerySessionDataResponseBody
	GetIndexData() []*QuerySessionDataResponseBodyIndexData
	SetMessage(v string) *QuerySessionDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySessionDataResponseBody
	GetRequestId() *string
}

type QuerySessionDataResponseBody struct {
	// The return code. This parameter is empty if the request is successful and data is found. If the query result is empty, NO_DATA is returned. If the request fails, the corresponding error code is returned. For example, if the format of the statistical parameter is invalid, InvalidParameter.DateFormat is returned.
	//
	// example:
	//
	// NO_DATA
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of session trend data, sorted in ascending order by hour and minute. If there is no session data on the specified date, an empty array is returned, and the return code is NO_DATA.
	IndexData []*QuerySessionDataResponseBodyIndexData `json:"IndexData,omitempty" xml:"IndexData,omitempty" type:"Repeated"`
	// The return message. This parameter is empty if the request is successful. If the request fails, the corresponding error description is returned.
	//
	// example:
	//
	// The parameter DateFormat is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, used to locate call records and troubleshoot issues. The example value is partially desensitized.
	//
	// example:
	//
	// D4A13D91-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySessionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySessionDataResponseBody) GetIndexData() []*QuerySessionDataResponseBodyIndexData {
	return s.IndexData
}

func (s *QuerySessionDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySessionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionDataResponseBody) SetCode(v string) *QuerySessionDataResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySessionDataResponseBody) SetIndexData(v []*QuerySessionDataResponseBodyIndexData) *QuerySessionDataResponseBody {
	s.IndexData = v
	return s
}

func (s *QuerySessionDataResponseBody) SetMessage(v string) *QuerySessionDataResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySessionDataResponseBody) SetRequestId(v string) *QuerySessionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionDataResponseBody) Validate() error {
	if s.IndexData != nil {
		for _, item := range s.IndexData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySessionDataResponseBodyIndexData struct {
	// The hour to which the session data belongs. Valid values: 00 to 23 (full hours).
	//
	// example:
	//
	// 10
	Hh *string `json:"Hh,omitempty" xml:"Hh,omitempty"`
	// The minute to which the session data belongs, aggregated in 10-minute intervals. Valid values: 00 to 50 (specifically 00, 10, 20, 30, 40, and 50).
	//
	// example:
	//
	// 30
	Mi *string `json:"Mi,omitempty" xml:"Mi,omitempty"`
	// The total number of sessions within the hour and minute interval (10-minute granularity).
	//
	// example:
	//
	// 128
	SessionCntMi *int64 `json:"SessionCntMi,omitempty" xml:"SessionCntMi,omitempty"`
	// The statistical date to which the session data belongs, in the yyyyMMdd format, which is the same as the StatDate request parameter.
	//
	// example:
	//
	// 20261106
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
}

func (s QuerySessionDataResponseBodyIndexData) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionDataResponseBodyIndexData) GoString() string {
	return s.String()
}

func (s *QuerySessionDataResponseBodyIndexData) GetHh() *string {
	return s.Hh
}

func (s *QuerySessionDataResponseBodyIndexData) GetMi() *string {
	return s.Mi
}

func (s *QuerySessionDataResponseBodyIndexData) GetSessionCntMi() *int64 {
	return s.SessionCntMi
}

func (s *QuerySessionDataResponseBodyIndexData) GetStatDate() *string {
	return s.StatDate
}

func (s *QuerySessionDataResponseBodyIndexData) SetHh(v string) *QuerySessionDataResponseBodyIndexData {
	s.Hh = &v
	return s
}

func (s *QuerySessionDataResponseBodyIndexData) SetMi(v string) *QuerySessionDataResponseBodyIndexData {
	s.Mi = &v
	return s
}

func (s *QuerySessionDataResponseBodyIndexData) SetSessionCntMi(v int64) *QuerySessionDataResponseBodyIndexData {
	s.SessionCntMi = &v
	return s
}

func (s *QuerySessionDataResponseBodyIndexData) SetStatDate(v string) *QuerySessionDataResponseBodyIndexData {
	s.StatDate = &v
	return s
}

func (s *QuerySessionDataResponseBodyIndexData) Validate() error {
	return dara.Validate(s)
}
