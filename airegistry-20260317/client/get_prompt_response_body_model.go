// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPromptResponseBodyData) *GetPromptResponseBody
	GetData() *GetPromptResponseBodyData
	SetRequestId(v string) *GetPromptResponseBody
	GetRequestId() *string
}

type GetPromptResponseBody struct {
	Data      *GetPromptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPromptResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBody) GetData() *GetPromptResponseBodyData {
	return s.Data
}

func (s *GetPromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPromptResponseBody) SetData(v *GetPromptResponseBodyData) *GetPromptResponseBody {
	s.Data = v
	return s
}

func (s *GetPromptResponseBody) SetRequestId(v string) *GetPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPromptResponseBodyData struct {
	BizTags          []*string                                  `json:"BizTags,omitempty" xml:"BizTags,omitempty" type:"Repeated"`
	Description      *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EditingVersion   *string                                    `json:"EditingVersion,omitempty" xml:"EditingVersion,omitempty"`
	GmtModified      *int64                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Labels           map[string]*string                         `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LatestVersion    *string                                    `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	OnlineCnt        *int32                                     `json:"OnlineCnt,omitempty" xml:"OnlineCnt,omitempty"`
	PromptKey        *string                                    `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	ReviewingVersion *string                                    `json:"ReviewingVersion,omitempty" xml:"ReviewingVersion,omitempty"`
	SchemaVersion    *int32                                     `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	VersionDetails   []*GetPromptResponseBodyDataVersionDetails `json:"VersionDetails,omitempty" xml:"VersionDetails,omitempty" type:"Repeated"`
	Versions         []*string                                  `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s GetPromptResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPromptResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBodyData) GetBizTags() []*string {
	return s.BizTags
}

func (s *GetPromptResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetPromptResponseBodyData) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *GetPromptResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetPromptResponseBodyData) GetLabels() map[string]*string {
	return s.Labels
}

func (s *GetPromptResponseBodyData) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetPromptResponseBodyData) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *GetPromptResponseBodyData) GetPromptKey() *string {
	return s.PromptKey
}

func (s *GetPromptResponseBodyData) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *GetPromptResponseBodyData) GetSchemaVersion() *int32 {
	return s.SchemaVersion
}

func (s *GetPromptResponseBodyData) GetVersionDetails() []*GetPromptResponseBodyDataVersionDetails {
	return s.VersionDetails
}

func (s *GetPromptResponseBodyData) GetVersions() []*string {
	return s.Versions
}

func (s *GetPromptResponseBodyData) SetBizTags(v []*string) *GetPromptResponseBodyData {
	s.BizTags = v
	return s
}

func (s *GetPromptResponseBodyData) SetDescription(v string) *GetPromptResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetPromptResponseBodyData) SetEditingVersion(v string) *GetPromptResponseBodyData {
	s.EditingVersion = &v
	return s
}

func (s *GetPromptResponseBodyData) SetGmtModified(v int64) *GetPromptResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetPromptResponseBodyData) SetLabels(v map[string]*string) *GetPromptResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetPromptResponseBodyData) SetLatestVersion(v string) *GetPromptResponseBodyData {
	s.LatestVersion = &v
	return s
}

func (s *GetPromptResponseBodyData) SetOnlineCnt(v int32) *GetPromptResponseBodyData {
	s.OnlineCnt = &v
	return s
}

func (s *GetPromptResponseBodyData) SetPromptKey(v string) *GetPromptResponseBodyData {
	s.PromptKey = &v
	return s
}

func (s *GetPromptResponseBodyData) SetReviewingVersion(v string) *GetPromptResponseBodyData {
	s.ReviewingVersion = &v
	return s
}

func (s *GetPromptResponseBodyData) SetSchemaVersion(v int32) *GetPromptResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *GetPromptResponseBodyData) SetVersionDetails(v []*GetPromptResponseBodyDataVersionDetails) *GetPromptResponseBodyData {
	s.VersionDetails = v
	return s
}

func (s *GetPromptResponseBodyData) SetVersions(v []*string) *GetPromptResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetPromptResponseBodyData) Validate() error {
	if s.VersionDetails != nil {
		for _, item := range s.VersionDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPromptResponseBodyDataVersionDetails struct {
	CommitMsg   *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	PromptKey   *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	SrcUser     *string `json:"SrcUser,omitempty" xml:"SrcUser,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPromptResponseBodyDataVersionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPromptResponseBodyDataVersionDetails) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBodyDataVersionDetails) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *GetPromptResponseBodyDataVersionDetails) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetPromptResponseBodyDataVersionDetails) GetPromptKey() *string {
	return s.PromptKey
}

func (s *GetPromptResponseBodyDataVersionDetails) GetSrcUser() *string {
	return s.SrcUser
}

func (s *GetPromptResponseBodyDataVersionDetails) GetStatus() *string {
	return s.Status
}

func (s *GetPromptResponseBodyDataVersionDetails) GetVersion() *string {
	return s.Version
}

func (s *GetPromptResponseBodyDataVersionDetails) SetCommitMsg(v string) *GetPromptResponseBodyDataVersionDetails {
	s.CommitMsg = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) SetGmtModified(v int64) *GetPromptResponseBodyDataVersionDetails {
	s.GmtModified = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) SetPromptKey(v string) *GetPromptResponseBodyDataVersionDetails {
	s.PromptKey = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) SetSrcUser(v string) *GetPromptResponseBodyDataVersionDetails {
	s.SrcUser = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) SetStatus(v string) *GetPromptResponseBodyDataVersionDetails {
	s.Status = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) SetVersion(v string) *GetPromptResponseBodyDataVersionDetails {
	s.Version = &v
	return s
}

func (s *GetPromptResponseBodyDataVersionDetails) Validate() error {
	return dara.Validate(s)
}
