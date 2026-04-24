// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySuccessIcpDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaSuccessDataWithRiskList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) *QuerySuccessIcpDataResponseBody
	GetBaSuccessDataWithRiskList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList
	SetErrorCode(v int32) *QuerySuccessIcpDataResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *QuerySuccessIcpDataResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QuerySuccessIcpDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySuccessIcpDataResponseBody
	GetSuccess() *bool
}

type QuerySuccessIcpDataResponseBody struct {
	BaSuccessDataWithRiskList []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList `json:"BaSuccessDataWithRiskList,omitempty" xml:"BaSuccessDataWithRiskList,omitempty" type:"Repeated"`
	// example:
	//
	// NoPermission
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1A13ABB5-7649-5031-B55C-D2E38F9F189D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySuccessIcpDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBody) GetBaSuccessDataWithRiskList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	return s.BaSuccessDataWithRiskList
}

func (s *QuerySuccessIcpDataResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *QuerySuccessIcpDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySuccessIcpDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySuccessIcpDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySuccessIcpDataResponseBody) SetBaSuccessDataWithRiskList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) *QuerySuccessIcpDataResponseBody {
	s.BaSuccessDataWithRiskList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBody) SetErrorCode(v int32) *QuerySuccessIcpDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBody) SetErrorMessage(v string) *QuerySuccessIcpDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBody) SetRequestId(v string) *QuerySuccessIcpDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBody) SetSuccess(v bool) *QuerySuccessIcpDataResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBody) Validate() error {
	if s.BaSuccessDataWithRiskList != nil {
		for _, item := range s.BaSuccessDataWithRiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList struct {
	AppList               []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList     `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	IcpNumber             *string                                                                `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
	OrganizersName        *string                                                                `json:"OrganizersName,omitempty" xml:"OrganizersName,omitempty"`
	OrganizersNature      *string                                                                `json:"OrganizersNature,omitempty" xml:"OrganizersNature,omitempty"`
	ResponsiblePersonName *string                                                                `json:"ResponsiblePersonName,omitempty" xml:"ResponsiblePersonName,omitempty"`
	RiskList              []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList    `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
	WebsiteList           []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList `json:"WebsiteList,omitempty" xml:"WebsiteList,omitempty" type:"Repeated"`
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetAppList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList {
	return s.AppList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetOrganizersName() *string {
	return s.OrganizersName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetOrganizersNature() *string {
	return s.OrganizersNature
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetResponsiblePersonName() *string {
	return s.ResponsiblePersonName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetRiskList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList {
	return s.RiskList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) GetWebsiteList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList {
	return s.WebsiteList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetAppList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.AppList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetIcpNumber(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.IcpNumber = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetOrganizersName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.OrganizersName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetOrganizersNature(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.OrganizersNature = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetResponsiblePersonName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.ResponsiblePersonName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetRiskList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.RiskList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) SetWebsiteList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList {
	s.WebsiteList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskList) Validate() error {
	if s.AppList != nil {
		for _, item := range s.AppList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebsiteList != nil {
		for _, item := range s.WebsiteList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList struct {
	// example:
	//
	// alipay
	AppName               *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppRecordNum          *string   `json:"AppRecordNum,omitempty" xml:"AppRecordNum,omitempty"`
	DomainList            []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	ResponsiblePersonName *string   `json:"ResponsiblePersonName,omitempty" xml:"ResponsiblePersonName,omitempty"`
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) GetAppName() *string {
	return s.AppName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) GetAppRecordNum() *string {
	return s.AppRecordNum
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) GetDomainList() []*string {
	return s.DomainList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) GetResponsiblePersonName() *string {
	return s.ResponsiblePersonName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) SetAppName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList {
	s.AppName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) SetAppRecordNum(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList {
	s.AppRecordNum = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) SetDomainList(v []*string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList {
	s.DomainList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) SetResponsiblePersonName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList {
	s.ResponsiblePersonName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListAppList) Validate() error {
	return dara.Validate(s)
}

type QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList struct {
	// example:
	//
	// 2026-04-24
	DeadLine       *string                                                                           `json:"DeadLine,omitempty" xml:"DeadLine,omitempty"`
	RiskDetailList []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList `json:"RiskDetailList,omitempty" xml:"RiskDetailList,omitempty" type:"Repeated"`
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) GetDeadLine() *string {
	return s.DeadLine
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) GetRiskDetailList() []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList {
	return s.RiskDetailList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) SetDeadLine(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList {
	s.DeadLine = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) SetRiskDetailList(v []*QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList {
	s.RiskDetailList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskList) Validate() error {
	if s.RiskDetailList != nil {
		for _, item := range s.RiskDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList struct {
	RectifySuggestList []*string `json:"RectifySuggestList,omitempty" xml:"RectifySuggestList,omitempty" type:"Repeated"`
	RiskSource         *string   `json:"RiskSource,omitempty" xml:"RiskSource,omitempty"`
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) GetRectifySuggestList() []*string {
	return s.RectifySuggestList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) GetRiskSource() *string {
	return s.RiskSource
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) SetRectifySuggestList(v []*string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList {
	s.RectifySuggestList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) SetRiskSource(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList {
	s.RiskSource = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListRiskListRiskDetailList) Validate() error {
	return dara.Validate(s)
}

type QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList struct {
	DomainList            []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	ResponsiblePersonName *string   `json:"ResponsiblePersonName,omitempty" xml:"ResponsiblePersonName,omitempty"`
	SiteName              *string   `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	SiteRecordNum         *string   `json:"SiteRecordNum,omitempty" xml:"SiteRecordNum,omitempty"`
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) GetDomainList() []*string {
	return s.DomainList
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) GetResponsiblePersonName() *string {
	return s.ResponsiblePersonName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) GetSiteName() *string {
	return s.SiteName
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) GetSiteRecordNum() *string {
	return s.SiteRecordNum
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) SetDomainList(v []*string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList {
	s.DomainList = v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) SetResponsiblePersonName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList {
	s.ResponsiblePersonName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) SetSiteName(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList {
	s.SiteName = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) SetSiteRecordNum(v string) *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList {
	s.SiteRecordNum = &v
	return s
}

func (s *QuerySuccessIcpDataResponseBodyBaSuccessDataWithRiskListWebsiteList) Validate() error {
	return dara.Validate(s)
}
