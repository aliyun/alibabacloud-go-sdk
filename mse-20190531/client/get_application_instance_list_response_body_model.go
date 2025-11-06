// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetApplicationInstanceListResponseBodyData) *GetApplicationInstanceListResponseBody
	GetData() *GetApplicationInstanceListResponseBodyData
}

type GetApplicationInstanceListResponseBody struct {
	// The returned data.
	Data *GetApplicationInstanceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetApplicationInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListResponseBody) GetData() *GetApplicationInstanceListResponseBodyData {
	return s.Data
}

func (s *GetApplicationInstanceListResponseBody) SetData(v *GetApplicationInstanceListResponseBodyData) *GetApplicationInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationInstanceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationInstanceListResponseBodyData struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried nodes.
	Result []*GetApplicationInstanceListResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of nodes.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetApplicationInstanceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetApplicationInstanceListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApplicationInstanceListResponseBodyData) GetResult() []*GetApplicationInstanceListResponseBodyDataResult {
	return s.Result
}

func (s *GetApplicationInstanceListResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetApplicationInstanceListResponseBodyData) SetPageNumber(v int32) *GetApplicationInstanceListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyData) SetPageSize(v int32) *GetApplicationInstanceListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyData) SetResult(v []*GetApplicationInstanceListResponseBodyDataResult) *GetApplicationInstanceListResponseBodyData {
	s.Result = v
	return s
}

func (s *GetApplicationInstanceListResponseBodyData) SetTotalSize(v int32) *GetApplicationInstanceListResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationInstanceListResponseBodyDataResult struct {
	// The node IP address.
	//
	// example:
	//
	// 10.1.2.3
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The application port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The node tags.
	Tags []*GetApplicationInstanceListResponseBodyDataResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetApplicationInstanceListResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListResponseBodyDataResult) GetIp() *string {
	return s.Ip
}

func (s *GetApplicationInstanceListResponseBodyDataResult) GetPort() *string {
	return s.Port
}

func (s *GetApplicationInstanceListResponseBodyDataResult) GetTags() []*GetApplicationInstanceListResponseBodyDataResultTags {
	return s.Tags
}

func (s *GetApplicationInstanceListResponseBodyDataResult) SetIp(v string) *GetApplicationInstanceListResponseBodyDataResult {
	s.Ip = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyDataResult) SetPort(v string) *GetApplicationInstanceListResponseBodyDataResult {
	s.Port = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyDataResult) SetTags(v []*GetApplicationInstanceListResponseBodyDataResultTags) *GetApplicationInstanceListResponseBodyDataResult {
	s.Tags = v
	return s
}

func (s *GetApplicationInstanceListResponseBodyDataResult) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationInstanceListResponseBodyDataResultTags struct {
	// 标签值。
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetApplicationInstanceListResponseBodyDataResultTags) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListResponseBodyDataResultTags) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListResponseBodyDataResultTags) GetTag() *string {
	return s.Tag
}

func (s *GetApplicationInstanceListResponseBodyDataResultTags) SetTag(v string) *GetApplicationInstanceListResponseBodyDataResultTags {
	s.Tag = &v
	return s
}

func (s *GetApplicationInstanceListResponseBodyDataResultTags) Validate() error {
	return dara.Validate(s)
}
