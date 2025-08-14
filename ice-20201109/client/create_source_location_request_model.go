// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseUrl(v string) *CreateSourceLocationRequest
	GetBaseUrl() *string
	SetEnableSegmentDelivery(v bool) *CreateSourceLocationRequest
	GetEnableSegmentDelivery() *bool
	SetSegmentDeliveryUrl(v string) *CreateSourceLocationRequest
	GetSegmentDeliveryUrl() *string
	SetSourceLocationName(v string) *CreateSourceLocationRequest
	GetSourceLocationName() *string
}

type CreateSourceLocationRequest struct {
	// The protocol and hostname of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://xxx.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// Specifies whether to use an independent domain name to access the segments.
	//
	// example:
	//
	// true
	EnableSegmentDelivery *bool `json:"EnableSegmentDelivery,omitempty" xml:"EnableSegmentDelivery,omitempty"`
	// The domain name used to access the segments.
	//
	// example:
	//
	// http://xxxxx.com
	SegmentDeliveryUrl *string `json:"SegmentDeliveryUrl,omitempty" xml:"SegmentDeliveryUrl,omitempty"`
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourcelocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
}

func (s CreateSourceLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceLocationRequest) GoString() string {
	return s.String()
}

func (s *CreateSourceLocationRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateSourceLocationRequest) GetEnableSegmentDelivery() *bool {
	return s.EnableSegmentDelivery
}

func (s *CreateSourceLocationRequest) GetSegmentDeliveryUrl() *string {
	return s.SegmentDeliveryUrl
}

func (s *CreateSourceLocationRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *CreateSourceLocationRequest) SetBaseUrl(v string) *CreateSourceLocationRequest {
	s.BaseUrl = &v
	return s
}

func (s *CreateSourceLocationRequest) SetEnableSegmentDelivery(v bool) *CreateSourceLocationRequest {
	s.EnableSegmentDelivery = &v
	return s
}

func (s *CreateSourceLocationRequest) SetSegmentDeliveryUrl(v string) *CreateSourceLocationRequest {
	s.SegmentDeliveryUrl = &v
	return s
}

func (s *CreateSourceLocationRequest) SetSourceLocationName(v string) *CreateSourceLocationRequest {
	s.SourceLocationName = &v
	return s
}

func (s *CreateSourceLocationRequest) Validate() error {
	return dara.Validate(s)
}
