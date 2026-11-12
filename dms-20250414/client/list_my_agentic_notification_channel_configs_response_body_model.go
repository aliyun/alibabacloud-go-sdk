// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAgenticNotificationChannelConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMyAgenticNotificationChannelConfigsResponseBodyData) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetData() []*ListMyAgenticNotificationChannelConfigsResponseBodyData
	SetErrorCode(v string) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMyAgenticNotificationChannelConfigsResponseBody
	GetSuccess() *bool
}

type ListMyAgenticNotificationChannelConfigsResponseBody struct {
	// The list of notification channel configurations for the current user.
	Data []*ListMyAgenticNotificationChannelConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// channel must be FEISHU
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of records per page used for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. This parameter is empty if no more data is available.
	//
	// example:
	//
	// MjA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMyAgenticNotificationChannelConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyAgenticNotificationChannelConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetData() []*ListMyAgenticNotificationChannelConfigsResponseBodyData {
	return s.Data
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetData(v []*ListMyAgenticNotificationChannelConfigsResponseBodyData) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetErrorCode(v string) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetErrorMessage(v string) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetMaxResults(v int32) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetNextToken(v string) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetRequestId(v string) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) SetSuccess(v bool) *ListMyAgenticNotificationChannelConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMyAgenticNotificationChannelConfigsResponseBodyData struct {
	// The notification channel type. Valid values: FEISHU (Lark).
	//
	// example:
	//
	// FEISHU
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The channel configuration information. Both keys and values are strings, which include bound identity information such as openId, mobile, or email. Sensitive values are always displayed as *****	- and are not returned in plaintext.
	//
	// example:
	//
	// {"openId":"******","mobile":"******","email":"******"}
	Config map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the notification channel is enabled. If the channel has never been configured, false is returned.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ListMyAgenticNotificationChannelConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMyAgenticNotificationChannelConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) GetChannel() *string {
	return s.Channel
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) GetConfig() map[string]*string {
	return s.Config
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) SetChannel(v string) *ListMyAgenticNotificationChannelConfigsResponseBodyData {
	s.Channel = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) SetConfig(v map[string]*string) *ListMyAgenticNotificationChannelConfigsResponseBodyData {
	s.Config = v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) SetEnabled(v bool) *ListMyAgenticNotificationChannelConfigsResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
