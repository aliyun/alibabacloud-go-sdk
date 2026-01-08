// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestWhatsappConversionApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *RequestWhatsappConversionApiRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *RequestWhatsappConversionApiRequest
	GetOwnerId() *int64
	SetPageId(v string) *RequestWhatsappConversionApiRequest
	GetPageId() *string
	SetRequestData(v []*RequestWhatsappConversionApiRequestRequestData) *RequestWhatsappConversionApiRequest
	GetRequestData() []*RequestWhatsappConversionApiRequestRequestData
	SetResourceOwnerAccount(v string) *RequestWhatsappConversionApiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RequestWhatsappConversionApiRequest
	GetResourceOwnerId() *int64
}

type RequestWhatsappConversionApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 929399382
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1939848838
	PageId               *string                                           `json:"PageId,omitempty" xml:"PageId,omitempty"`
	RequestData          []*RequestWhatsappConversionApiRequestRequestData `json:"RequestData,omitempty" xml:"RequestData,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RequestWhatsappConversionApiRequest) String() string {
	return dara.Prettify(s)
}

func (s RequestWhatsappConversionApiRequest) GoString() string {
	return s.String()
}

func (s *RequestWhatsappConversionApiRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *RequestWhatsappConversionApiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RequestWhatsappConversionApiRequest) GetPageId() *string {
	return s.PageId
}

func (s *RequestWhatsappConversionApiRequest) GetRequestData() []*RequestWhatsappConversionApiRequestRequestData {
	return s.RequestData
}

func (s *RequestWhatsappConversionApiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RequestWhatsappConversionApiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RequestWhatsappConversionApiRequest) SetCustSpaceId(v string) *RequestWhatsappConversionApiRequest {
	s.CustSpaceId = &v
	return s
}

func (s *RequestWhatsappConversionApiRequest) SetOwnerId(v int64) *RequestWhatsappConversionApiRequest {
	s.OwnerId = &v
	return s
}

func (s *RequestWhatsappConversionApiRequest) SetPageId(v string) *RequestWhatsappConversionApiRequest {
	s.PageId = &v
	return s
}

func (s *RequestWhatsappConversionApiRequest) SetRequestData(v []*RequestWhatsappConversionApiRequestRequestData) *RequestWhatsappConversionApiRequest {
	s.RequestData = v
	return s
}

func (s *RequestWhatsappConversionApiRequest) SetResourceOwnerAccount(v string) *RequestWhatsappConversionApiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RequestWhatsappConversionApiRequest) SetResourceOwnerId(v int64) *RequestWhatsappConversionApiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RequestWhatsappConversionApiRequest) Validate() error {
	if s.RequestData != nil {
		for _, item := range s.RequestData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RequestWhatsappConversionApiRequestRequestData struct {
	// This parameter is required.
	//
	// example:
	//
	// business_messaging
	ActionSource          *string                `json:"ActionSource,omitempty" xml:"ActionSource,omitempty"`
	AppData               map[string]interface{} `json:"AppData,omitempty" xml:"AppData,omitempty"`
	CustomData            map[string]interface{} `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
	DataProcessingOptions []*string              `json:"DataProcessingOptions,omitempty" xml:"DataProcessingOptions,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	DataProcessingOptionsCountry *int64 `json:"DataProcessingOptionsCountry,omitempty" xml:"DataProcessingOptionsCountry,omitempty"`
	// example:
	//
	// 26
	DataProcessingOptionsState *int64 `json:"DataProcessingOptionsState,omitempty" xml:"DataProcessingOptionsState,omitempty"`
	// example:
	//
	// 20029399299
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Purchase
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// http://alibaba.com
	EventSourceUrl *string `json:"EventSourceUrl,omitempty" xml:"EventSourceUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1709201870
	EventTime *int64                 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	ExtInfo   map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// whatsapp
	MessagingChannel *string `json:"MessagingChannel,omitempty" xml:"MessagingChannel,omitempty"`
	// example:
	//
	// true
	OptOut *bool `json:"OptOut,omitempty" xml:"OptOut,omitempty"`
	// This parameter is required.
	UserData map[string]interface{} `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s RequestWhatsappConversionApiRequestRequestData) String() string {
	return dara.Prettify(s)
}

func (s RequestWhatsappConversionApiRequestRequestData) GoString() string {
	return s.String()
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetActionSource() *string {
	return s.ActionSource
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetAppData() map[string]interface{} {
	return s.AppData
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetCustomData() map[string]interface{} {
	return s.CustomData
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetDataProcessingOptions() []*string {
	return s.DataProcessingOptions
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetDataProcessingOptionsCountry() *int64 {
	return s.DataProcessingOptionsCountry
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetDataProcessingOptionsState() *int64 {
	return s.DataProcessingOptionsState
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetEventId() *string {
	return s.EventId
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetEventName() *string {
	return s.EventName
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetEventSourceUrl() *string {
	return s.EventSourceUrl
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetEventTime() *int64 {
	return s.EventTime
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetMessagingChannel() *string {
	return s.MessagingChannel
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetOptOut() *bool {
	return s.OptOut
}

func (s *RequestWhatsappConversionApiRequestRequestData) GetUserData() map[string]interface{} {
	return s.UserData
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetActionSource(v string) *RequestWhatsappConversionApiRequestRequestData {
	s.ActionSource = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetAppData(v map[string]interface{}) *RequestWhatsappConversionApiRequestRequestData {
	s.AppData = v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetCustomData(v map[string]interface{}) *RequestWhatsappConversionApiRequestRequestData {
	s.CustomData = v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetDataProcessingOptions(v []*string) *RequestWhatsappConversionApiRequestRequestData {
	s.DataProcessingOptions = v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetDataProcessingOptionsCountry(v int64) *RequestWhatsappConversionApiRequestRequestData {
	s.DataProcessingOptionsCountry = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetDataProcessingOptionsState(v int64) *RequestWhatsappConversionApiRequestRequestData {
	s.DataProcessingOptionsState = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetEventId(v string) *RequestWhatsappConversionApiRequestRequestData {
	s.EventId = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetEventName(v string) *RequestWhatsappConversionApiRequestRequestData {
	s.EventName = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetEventSourceUrl(v string) *RequestWhatsappConversionApiRequestRequestData {
	s.EventSourceUrl = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetEventTime(v int64) *RequestWhatsappConversionApiRequestRequestData {
	s.EventTime = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetExtInfo(v map[string]interface{}) *RequestWhatsappConversionApiRequestRequestData {
	s.ExtInfo = v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetMessagingChannel(v string) *RequestWhatsappConversionApiRequestRequestData {
	s.MessagingChannel = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetOptOut(v bool) *RequestWhatsappConversionApiRequestRequestData {
	s.OptOut = &v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) SetUserData(v map[string]interface{}) *RequestWhatsappConversionApiRequestRequestData {
	s.UserData = v
	return s
}

func (s *RequestWhatsappConversionApiRequestRequestData) Validate() error {
	return dara.Validate(s)
}
