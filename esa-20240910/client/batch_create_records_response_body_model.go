// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordResultList(v *BatchCreateRecordsResponseBodyRecordResultList) *BatchCreateRecordsResponseBody
	GetRecordResultList() *BatchCreateRecordsResponseBodyRecordResultList
	SetRequestId(v string) *BatchCreateRecordsResponseBody
	GetRequestId() *string
}

type BatchCreateRecordsResponseBody struct {
	RecordResultList *BatchCreateRecordsResponseBodyRecordResultList `json:"RecordResultList,omitempty" xml:"RecordResultList,omitempty" type:"Struct"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBody) GetRecordResultList() *BatchCreateRecordsResponseBodyRecordResultList {
	return s.RecordResultList
}

func (s *BatchCreateRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateRecordsResponseBody) SetRecordResultList(v *BatchCreateRecordsResponseBodyRecordResultList) *BatchCreateRecordsResponseBody {
	s.RecordResultList = v
	return s
}

func (s *BatchCreateRecordsResponseBody) SetRequestId(v string) *BatchCreateRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateRecordsResponseBody) Validate() error {
	if s.RecordResultList != nil {
		if err := s.RecordResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultList struct {
	Failed     []*BatchCreateRecordsResponseBodyRecordResultListFailed  `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Repeated"`
	Success    []*BatchCreateRecordsResponseBodyRecordResultListSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultList) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetFailed() []*BatchCreateRecordsResponseBodyRecordResultListFailed {
	return s.Failed
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetSuccess() []*BatchCreateRecordsResponseBodyRecordResultListSuccess {
	return s.Success
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetFailed(v []*BatchCreateRecordsResponseBodyRecordResultListFailed) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Failed = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetSuccess(v []*BatchCreateRecordsResponseBodyRecordResultListSuccess) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Success = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetTotalCount(v int32) *BatchCreateRecordsResponseBodyRecordResultList {
	s.TotalCount = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) Validate() error {
	if s.Failed != nil {
		for _, item := range s.Failed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Success != nil {
		for _, item := range s.Success {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListFailed struct {
	BizName     *string                                                   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Data        *BatchCreateRecordsResponseBodyRecordResultListFailedData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Description *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Proxied     *bool                                                     `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	RecordId    *int64                                                    `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordName  *string                                                   `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	RecordType  *string                                                   `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SourceType  *string                                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl         *int32                                                    `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetBizName() *string {
	return s.BizName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetData() *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	return s.Data
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetProxied() *bool {
	return s.Proxied
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordId() *int64 {
	return s.RecordId
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordName() *string {
	return s.RecordName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordType() *string {
	return s.RecordType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetTtl() *int32 {
	return s.Ttl
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetData(v *BatchCreateRecordsResponseBodyRecordResultListFailedData) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListFailedData struct {
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Fingerprint  *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Flag         *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	MatchingType *int32  `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Selector     *int32  `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Usage        *int32  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetCertificate() *string {
	return s.Certificate
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetFlag() *int32 {
	return s.Flag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetPort() *int32 {
	return s.Port
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetSelector() *int32 {
	return s.Selector
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetTag() *string {
	return s.Tag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetType() *int32 {
	return s.Type
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetUsage() *int32 {
	return s.Usage
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetValue() *string {
	return s.Value
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetWeight() *int32 {
	return s.Weight
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Weight = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultListSuccess struct {
	BizName     *string                                                    `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Data        *BatchCreateRecordsResponseBodyRecordResultListSuccessData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Description *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Proxied     *bool                                                      `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	RecordId    *int64                                                     `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordName  *string                                                    `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	RecordType  *string                                                    `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SourceType  *string                                                    `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl         *int32                                                     `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetBizName() *string {
	return s.BizName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetData() *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	return s.Data
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetProxied() *bool {
	return s.Proxied
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordId() *int64 {
	return s.RecordId
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordName() *string {
	return s.RecordName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordType() *string {
	return s.RecordType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetTtl() *int32 {
	return s.Ttl
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetData(v *BatchCreateRecordsResponseBodyRecordResultListSuccessData) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListSuccessData struct {
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Fingerprint  *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Flag         *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	MatchingType *int32  `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Selector     *int32  `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Usage        *int32  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetCertificate() *string {
	return s.Certificate
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetFlag() *int32 {
	return s.Flag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetPort() *int32 {
	return s.Port
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetSelector() *int32 {
	return s.Selector
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetTag() *string {
	return s.Tag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetType() *int32 {
	return s.Type
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetUsage() *int32 {
	return s.Usage
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetValue() *string {
	return s.Value
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetWeight() *int32 {
	return s.Weight
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Weight = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) Validate() error {
	return dara.Validate(s)
}
