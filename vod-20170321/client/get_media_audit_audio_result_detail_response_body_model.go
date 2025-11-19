// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditAudioResultDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditAudioResultDetail(v *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) *GetMediaAuditAudioResultDetailResponseBody
	GetMediaAuditAudioResultDetail() *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail
	SetRequestId(v string) *GetMediaAuditAudioResultDetailResponseBody
	GetRequestId() *string
}

type GetMediaAuditAudioResultDetailResponseBody struct {
	// Details of review results.
	MediaAuditAudioResultDetail *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail `json:"MediaAuditAudioResultDetail,omitempty" xml:"MediaAuditAudioResultDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CB7D7232-1AB2-40FE-B8D3-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBody) GetMediaAuditAudioResultDetail() *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	return s.MediaAuditAudioResultDetail
}

func (s *GetMediaAuditAudioResultDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaAuditAudioResultDetailResponseBody) SetMediaAuditAudioResultDetail(v *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) *GetMediaAuditAudioResultDetailResponseBody {
	s.MediaAuditAudioResultDetail = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBody) SetRequestId(v string) *GetMediaAuditAudioResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBody) Validate() error {
	if s.MediaAuditAudioResultDetail != nil {
		if err := s.MediaAuditAudioResultDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail struct {
	// The list of results.
	List []*GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) GetList() []*GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	return s.List
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) GetPageTotal() *int32 {
	return s.PageTotal
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) GetTotal() *int32 {
	return s.Total
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetList(v []*GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.List = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetPageTotal(v int32) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.PageTotal = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetTotal(v int32) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.Total = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList struct {
	// The end time of the audio that failed the review. Unit: seconds.
	//
	// example:
	//
	// 10
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The review results. Valid values:
	//
	// 	- **spam**
	//
	// 	- **ad**
	//
	// 	- **abuse**
	//
	// 	- **flood**
	//
	// 	- **contraband**
	//
	// 	- **meaningless**
	//
	// 	- **normal**
	//
	// example:
	//
	// abuse
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The start time of the audio that failed the review. Unit: seconds.
	//
	// example:
	//
	// 8
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The text that corresponds to the audio.
	//
	// example:
	//
	// beauty
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GetText() *string {
	return s.Text
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetEndTime(v int64) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.EndTime = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetLabel(v string) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetStartTime(v int64) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.StartTime = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetText(v string) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.Text = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) Validate() error {
	return dara.Validate(s)
}
