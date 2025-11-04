// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectAvailableRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *GetMediaConnectAvailableRegionResponseBodyContent) *GetMediaConnectAvailableRegionResponseBody
	GetContent() *GetMediaConnectAvailableRegionResponseBodyContent
	SetDescription(v string) *GetMediaConnectAvailableRegionResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetMediaConnectAvailableRegionResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *GetMediaConnectAvailableRegionResponseBody
	GetRetCode() *int32
}

type GetMediaConnectAvailableRegionResponseBody struct {
	// The rsponse body.
	Content *GetMediaConnectAvailableRegionResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The call description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 41CB9D4C-4650-5723-BA89-D6824F706ACB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s GetMediaConnectAvailableRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectAvailableRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConnectAvailableRegionResponseBody) GetContent() *GetMediaConnectAvailableRegionResponseBodyContent {
	return s.Content
}

func (s *GetMediaConnectAvailableRegionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMediaConnectAvailableRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConnectAvailableRegionResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetMediaConnectAvailableRegionResponseBody) SetContent(v *GetMediaConnectAvailableRegionResponseBodyContent) *GetMediaConnectAvailableRegionResponseBody {
	s.Content = v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBody) SetDescription(v string) *GetMediaConnectAvailableRegionResponseBody {
	s.Description = &v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBody) SetRequestId(v string) *GetMediaConnectAvailableRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBody) SetRetCode(v int32) *GetMediaConnectAvailableRegionResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaConnectAvailableRegionResponseBodyContent struct {
	// The default region. You can ignore the parameter.
	//
	// example:
	//
	// cn-shanghai
	DefaultRegion *string `json:"DefaultRegion,omitempty" xml:"DefaultRegion,omitempty"`
	// The supported regions.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s GetMediaConnectAvailableRegionResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectAvailableRegionResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetMediaConnectAvailableRegionResponseBodyContent) GetDefaultRegion() *string {
	return s.DefaultRegion
}

func (s *GetMediaConnectAvailableRegionResponseBodyContent) GetRegionList() []*string {
	return s.RegionList
}

func (s *GetMediaConnectAvailableRegionResponseBodyContent) SetDefaultRegion(v string) *GetMediaConnectAvailableRegionResponseBodyContent {
	s.DefaultRegion = &v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBodyContent) SetRegionList(v []*string) *GetMediaConnectAvailableRegionResponseBodyContent {
	s.RegionList = v
	return s
}

func (s *GetMediaConnectAvailableRegionResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
