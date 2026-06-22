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
	// The space ID or instance ID of the customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 929399382
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PageId of Meta.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1939848838
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The request data.
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
	// Specifies where the conversion occurred. Knowing where the event occurred helps ensure that ads are delivered to the correct audience. By using the Conversions API, you agree that the action_source parameter is accurate to the best of your knowledge.
	//
	//
	// The values you can send in the action_source field are as follows:
	//
	//
	// - email: The conversion occurred through email.
	//
	// - website: The conversion was made on a website.
	//
	// - app: The conversion was made on a shift application.
	//
	// - phone_call: The conversion was made over the phone.
	//
	// - chat: The conversion was made through a messaging app, SMS, or online messaging feature.
	//
	// - physical_store: The conversion was made in person at a physical store entity.
	//
	// - system_generated: The conversion occurred automatically, such as a subscribe renewal with Settings for monthly automatic payment.
	//
	// - other: The conversion was made through a method not listed in this topic.
	//
	// Note: All action source values support ads measurement and custom audience creation. All action sources except physical_store support ads optimization.
	//
	// This parameter is required.
	//
	// example:
	//
	// business_messaging
	ActionSource *string `json:"ActionSource,omitempty" xml:"ActionSource,omitempty"`
	// Required parameters for app events.
	//
	// These parameters are used to share app data and device information with the Conversions API.
	AppData map[string]interface{} `json:"AppData,omitempty" xml:"AppData,omitempty"`
	// A map that contains additional business data for the event.
	CustomData map[string]interface{} `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
	// The processing options you want to enable for a specific event. For Limited Data Use, the currently accepted value is LDU. You can send an empty array to explicitly specify that the event must not be processed with Limited Data Use restrictions.
	DataProcessingOptions []*string `json:"DataProcessingOptions,omitempty" xml:"DataProcessingOptions,omitempty" type:"Repeated"`
	// Required if you send LDU under data_processing_options.
	//
	// The country you want to associate with this data processing option. Currently accepted values are 1 (representing the United States) or 0 (requesting that we geolocate this event).
	//
	// example:
	//
	// 100
	DataProcessingOptionsCountry *int64 `json:"DataProcessingOptionsCountry,omitempty" xml:"DataProcessingOptionsCountry,omitempty"`
	// Required in some cases. (See the notes below for details.)
	//
	// The state you want to associate with this data processing option. Currently accepted values are 1000 (representing California) or 0 (requesting that we geolocate this event).
	//
	// example:
	//
	// 26
	DataProcessingOptionsState *int64 `json:"DataProcessingOptionsState,omitempty" xml:"DataProcessingOptionsState,omitempty"`
	// This ID can be any unique string chosen by the advertiser. The event_name and event_id parameters are used to deduplicate events sent by a website (through Meta Pixel), an app (through the SDK or App Events API), and the Conversions API. Although event_id is marked as optional, we recommend that you use this parameter for deduplication.
	//
	// example:
	//
	// 20029399299
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// - The name of a standard event or custom event. This field is used to deduplicate events sent by a website (through Meta Pixel), an app (through the SDK or App Events API), and the Conversions API. The event_id parameter is also used for deduplication.
	//
	// - For the same customer action, the event from the browser or app should match the event_name from the server event. If a match is found between events sent within 48 hours, only the first event is considered. If server and browser/app events are received at approximately the same time (within 5 minutes of each other), the browser/app event takes priority. Learn more about deduplicating Pixel events and server events.
	//
	// This parameter is required.
	//
	// example:
	//
	// Purchase
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The browser URL where the event occurred. The URL must start with http:// or https:// and should match the verified domain.
	//
	// example:
	//
	// http://alibaba.com
	EventSourceUrl *string `json:"EventSourceUrl,omitempty" xml:"EventSourceUrl,omitempty"`
	// A Unix timestamp in seconds indicating when the event actually occurred. The specified time may be earlier than the time you send the event to Facebook. This is intended for batch processing and server performance optimization. You must send the date in Greenwich Mean Time (GMT) time zone format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1709201870
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// Required parameters for app events.
	//
	// Extended device information, such as the width and height of the screen. This parameter is an array with values separated by commas. When using extinfo, all values are required and must be arranged in the following index order. If a value is missing, use an empty string as a placeholder.
	ExtInfo map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// The source. Fixed value: whatsapp.
	//
	// example:
	//
	// whatsapp
	MessagingChannel *string `json:"MessagingChannel,omitempty" xml:"MessagingChannel,omitempty"`
	// A flag that indicates this event should not be used for ad delivery optimization. When set to true, the event can only be used for attribution.
	//
	// example:
	//
	// true
	OptOut *bool `json:"OptOut,omitempty" xml:"OptOut,omitempty"`
	// A map that contains customer information data.
	//
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
