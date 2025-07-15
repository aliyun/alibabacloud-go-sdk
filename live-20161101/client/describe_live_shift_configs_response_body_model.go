// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveShiftConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeLiveShiftConfigsResponseBodyContent) *DescribeLiveShiftConfigsResponseBody
	GetContent() *DescribeLiveShiftConfigsResponseBodyContent
	SetRequestId(v string) *DescribeLiveShiftConfigsResponseBody
	GetRequestId() *string
}

type DescribeLiveShiftConfigsResponseBody struct {
	// The time shifting configurations.
	Content *DescribeLiveShiftConfigsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveShiftConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveShiftConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveShiftConfigsResponseBody) GetContent() *DescribeLiveShiftConfigsResponseBodyContent {
	return s.Content
}

func (s *DescribeLiveShiftConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveShiftConfigsResponseBody) SetContent(v *DescribeLiveShiftConfigsResponseBodyContent) *DescribeLiveShiftConfigsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBody) SetRequestId(v string) *DescribeLiveShiftConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveShiftConfigsResponseBodyContent struct {
	Config []*DescribeLiveShiftConfigsResponseBodyContentConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s DescribeLiveShiftConfigsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveShiftConfigsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeLiveShiftConfigsResponseBodyContent) GetConfig() []*DescribeLiveShiftConfigsResponseBodyContentConfig {
	return s.Config
}

func (s *DescribeLiveShiftConfigsResponseBodyContent) SetConfig(v []*DescribeLiveShiftConfigsResponseBodyContentConfig) *DescribeLiveShiftConfigsResponseBodyContent {
	s.Config = v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveShiftConfigsResponseBodyContentConfig struct {
	// The application for which you configure time shifting.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name for which you configure time shifting.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Whether to ignore time shift generation for the transcode stream.
	//
	// 	- true: Ignore time shifting generation.
	//
	// 	- false: Generate time shifting.
	//
	// The default value is true.
	//
	// example:
	//
	// true
	IgnoreTranscode *bool `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	// The name of the live stream for which you configure time shifting.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The number of days for which the time shifting configurations are retained.
	//
	// example:
	//
	// 7
	Vision *int32 `json:"Vision,omitempty" xml:"Vision,omitempty"`
}

func (s DescribeLiveShiftConfigsResponseBodyContentConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveShiftConfigsResponseBodyContentConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) GetIgnoreTranscode() *bool {
	return s.IgnoreTranscode
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) GetVision() *int32 {
	return s.Vision
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) SetAppName(v string) *DescribeLiveShiftConfigsResponseBodyContentConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) SetDomainName(v string) *DescribeLiveShiftConfigsResponseBodyContentConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) SetIgnoreTranscode(v bool) *DescribeLiveShiftConfigsResponseBodyContentConfig {
	s.IgnoreTranscode = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) SetStreamName(v string) *DescribeLiveShiftConfigsResponseBodyContentConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) SetVision(v int32) *DescribeLiveShiftConfigsResponseBodyContentConfig {
	s.Vision = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponseBodyContentConfig) Validate() error {
	return dara.Validate(s)
}
