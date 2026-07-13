// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFeaturesResponseBody
	GetCode() *string
	SetData(v *QueryFeaturesResponseBodyData) *QueryFeaturesResponseBody
	GetData() *QueryFeaturesResponseBodyData
	SetHttpStatusCode(v int32) *QueryFeaturesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryFeaturesResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryFeaturesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFeaturesResponseBody
	GetSuccess() *bool
}

type QueryFeaturesResponseBody struct {
	// example:
	//
	// Success
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryFeaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFeaturesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFeaturesResponseBody) GetData() *QueryFeaturesResponseBodyData {
	return s.Data
}

func (s *QueryFeaturesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryFeaturesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFeaturesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFeaturesResponseBody) SetCode(v string) *QueryFeaturesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFeaturesResponseBody) SetData(v *QueryFeaturesResponseBodyData) *QueryFeaturesResponseBody {
	s.Data = v
	return s
}

func (s *QueryFeaturesResponseBody) SetHttpStatusCode(v int32) *QueryFeaturesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryFeaturesResponseBody) SetMessage(v string) *QueryFeaturesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFeaturesResponseBody) SetRequestId(v string) *QueryFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFeaturesResponseBody) SetSuccess(v bool) *QueryFeaturesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFeaturesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFeaturesResponseBodyData struct {
	Features     []*QueryFeaturesResponseBodyDataFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	InstanceId   *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceName *string                                  `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	TargetScope  *string                                  `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s QueryFeaturesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFeaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFeaturesResponseBodyData) GetFeatures() []*QueryFeaturesResponseBodyDataFeatures {
	return s.Features
}

func (s *QueryFeaturesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryFeaturesResponseBodyData) GetResourceName() *string {
	return s.ResourceName
}

func (s *QueryFeaturesResponseBodyData) GetTargetScope() *string {
	return s.TargetScope
}

func (s *QueryFeaturesResponseBodyData) SetFeatures(v []*QueryFeaturesResponseBodyDataFeatures) *QueryFeaturesResponseBodyData {
	s.Features = v
	return s
}

func (s *QueryFeaturesResponseBodyData) SetInstanceId(v string) *QueryFeaturesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryFeaturesResponseBodyData) SetResourceName(v string) *QueryFeaturesResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *QueryFeaturesResponseBodyData) SetTargetScope(v string) *QueryFeaturesResponseBodyData {
	s.TargetScope = &v
	return s
}

func (s *QueryFeaturesResponseBodyData) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFeaturesResponseBodyDataFeatures struct {
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName           *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FeatureCode           *string `json:"FeatureCode,omitempty" xml:"FeatureCode,omitempty"`
	Supported             *bool   `json:"Supported,omitempty" xml:"Supported,omitempty"`
	UnsupportedReason     *string `json:"UnsupportedReason,omitempty" xml:"UnsupportedReason,omitempty"`
	UnsupportedReasonCode *string `json:"UnsupportedReasonCode,omitempty" xml:"UnsupportedReasonCode,omitempty"`
}

func (s QueryFeaturesResponseBodyDataFeatures) String() string {
	return dara.Prettify(s)
}

func (s QueryFeaturesResponseBodyDataFeatures) GoString() string {
	return s.String()
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetDescription() *string {
	return s.Description
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetDisplayName() *string {
	return s.DisplayName
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetFeatureCode() *string {
	return s.FeatureCode
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetSupported() *bool {
	return s.Supported
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetUnsupportedReason() *string {
	return s.UnsupportedReason
}

func (s *QueryFeaturesResponseBodyDataFeatures) GetUnsupportedReasonCode() *string {
	return s.UnsupportedReasonCode
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetDescription(v string) *QueryFeaturesResponseBodyDataFeatures {
	s.Description = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetDisplayName(v string) *QueryFeaturesResponseBodyDataFeatures {
	s.DisplayName = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetFeatureCode(v string) *QueryFeaturesResponseBodyDataFeatures {
	s.FeatureCode = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetSupported(v bool) *QueryFeaturesResponseBodyDataFeatures {
	s.Supported = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetUnsupportedReason(v string) *QueryFeaturesResponseBodyDataFeatures {
	s.UnsupportedReason = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) SetUnsupportedReasonCode(v string) *QueryFeaturesResponseBodyDataFeatures {
	s.UnsupportedReasonCode = &v
	return s
}

func (s *QueryFeaturesResponseBodyDataFeatures) Validate() error {
	return dara.Validate(s)
}
