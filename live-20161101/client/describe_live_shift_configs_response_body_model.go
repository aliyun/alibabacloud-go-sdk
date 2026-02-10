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
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Config != nil {
		for _, item := range s.Config {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveShiftConfigsResponseBodyContentConfig struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IgnoreTranscode *bool   `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	StreamName      *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Vision          *int32  `json:"Vision,omitempty" xml:"Vision,omitempty"`
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
