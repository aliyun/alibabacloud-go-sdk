// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncEcsHostTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSyncEcsHostTaskRequest
	GetLang() *string
	SetZoneId(v string) *DescribeSyncEcsHostTaskRequest
	GetZoneId() *string
}

type DescribeSyncEcsHostTaskRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// pvtz-test-id-2989149d628c5****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSyncEcsHostTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSyncEcsHostTaskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSyncEcsHostTaskRequest) SetLang(v string) *DescribeSyncEcsHostTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSyncEcsHostTaskRequest) SetZoneId(v string) *DescribeSyncEcsHostTaskRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSyncEcsHostTaskRequest) Validate() error {
	return dara.Validate(s)
}
