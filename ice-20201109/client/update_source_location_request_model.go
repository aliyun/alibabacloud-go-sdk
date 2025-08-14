// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseUrl(v string) *UpdateSourceLocationRequest
	GetBaseUrl() *string
	SetEnableSegmentDelivery(v bool) *UpdateSourceLocationRequest
	GetEnableSegmentDelivery() *bool
	SetSegmentDeliveryUrl(v string) *UpdateSourceLocationRequest
	GetSegmentDeliveryUrl() *string
	SetSourceLocationName(v string) *UpdateSourceLocationRequest
	GetSourceLocationName() *string
}

type UpdateSourceLocationRequest struct {
	// The protocol and hostname of the source location.
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
	// http://xxxx.com
	SegmentDeliveryUrl *string `json:"SegmentDeliveryUrl,omitempty" xml:"SegmentDeliveryUrl,omitempty"`
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
}

func (s UpdateSourceLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceLocationRequest) GoString() string {
	return s.String()
}

func (s *UpdateSourceLocationRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *UpdateSourceLocationRequest) GetEnableSegmentDelivery() *bool {
	return s.EnableSegmentDelivery
}

func (s *UpdateSourceLocationRequest) GetSegmentDeliveryUrl() *string {
	return s.SegmentDeliveryUrl
}

func (s *UpdateSourceLocationRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *UpdateSourceLocationRequest) SetBaseUrl(v string) *UpdateSourceLocationRequest {
	s.BaseUrl = &v
	return s
}

func (s *UpdateSourceLocationRequest) SetEnableSegmentDelivery(v bool) *UpdateSourceLocationRequest {
	s.EnableSegmentDelivery = &v
	return s
}

func (s *UpdateSourceLocationRequest) SetSegmentDeliveryUrl(v string) *UpdateSourceLocationRequest {
	s.SegmentDeliveryUrl = &v
	return s
}

func (s *UpdateSourceLocationRequest) SetSourceLocationName(v string) *UpdateSourceLocationRequest {
	s.SourceLocationName = &v
	return s
}

func (s *UpdateSourceLocationRequest) Validate() error {
	return dara.Validate(s)
}
