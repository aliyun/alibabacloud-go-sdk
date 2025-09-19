// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLevelsOn(v []*string) *DescribeSuspEventUserSettingResponseBody
	GetLevelsOn() []*string
	SetRequestId(v string) *DescribeSuspEventUserSettingResponseBody
	GetRequestId() *string
}

type DescribeSuspEventUserSettingResponseBody struct {
	// An array that consists of the risk levels of alert notifications. Valid values:
	//
	// 	- **remind**
	//
	// 	- **suspicious**
	//
	// 	- **serious**
	LevelsOn []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSuspEventUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventUserSettingResponseBody) GetLevelsOn() []*string {
	return s.LevelsOn
}

func (s *DescribeSuspEventUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventUserSettingResponseBody) SetLevelsOn(v []*string) *DescribeSuspEventUserSettingResponseBody {
	s.LevelsOn = v
	return s
}

func (s *DescribeSuspEventUserSettingResponseBody) SetRequestId(v string) *DescribeSuspEventUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
