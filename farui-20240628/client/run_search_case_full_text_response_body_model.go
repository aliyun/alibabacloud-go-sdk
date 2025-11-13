// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchCaseFullTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunSearchCaseFullTextResponseBody
	GetCode() *string
	SetData(v *RunSearchCaseFullTextResponseBodyData) *RunSearchCaseFullTextResponseBody
	GetData() *RunSearchCaseFullTextResponseBodyData
	SetHttpStatusCode(v int64) *RunSearchCaseFullTextResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RunSearchCaseFullTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunSearchCaseFullTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunSearchCaseFullTextResponseBody
	GetSuccess() *bool
}

type RunSearchCaseFullTextResponseBody struct {
	// example:
	//
	// null
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *RunSearchCaseFullTextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C844BE6B-33A9-5AC4-A1AE-97B131849E0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunSearchCaseFullTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunSearchCaseFullTextResponseBody) GetData() *RunSearchCaseFullTextResponseBodyData {
	return s.Data
}

func (s *RunSearchCaseFullTextResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RunSearchCaseFullTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunSearchCaseFullTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSearchCaseFullTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunSearchCaseFullTextResponseBody) SetCode(v string) *RunSearchCaseFullTextResponseBody {
	s.Code = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetData(v *RunSearchCaseFullTextResponseBodyData) *RunSearchCaseFullTextResponseBody {
	s.Data = v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetHttpStatusCode(v int64) *RunSearchCaseFullTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetMessage(v string) *RunSearchCaseFullTextResponseBody {
	s.Message = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetRequestId(v string) *RunSearchCaseFullTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetSuccess(v bool) *RunSearchCaseFullTextResponseBody {
	s.Success = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchCaseFullTextResponseBodyData struct {
	CaseLevel  *string                                            `json:"caseLevel,omitempty" xml:"caseLevel,omitempty"`
	CaseResult []*RunSearchCaseFullTextResponseBodyDataCaseResult `json:"caseResult,omitempty" xml:"caseResult,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query         *string   `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords []*string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyData) GetCaseLevel() *string {
	return s.CaseLevel
}

func (s *RunSearchCaseFullTextResponseBodyData) GetCaseResult() []*RunSearchCaseFullTextResponseBodyDataCaseResult {
	return s.CaseResult
}

func (s *RunSearchCaseFullTextResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *RunSearchCaseFullTextResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunSearchCaseFullTextResponseBodyData) GetQuery() *string {
	return s.Query
}

func (s *RunSearchCaseFullTextResponseBodyData) GetQueryKeywords() []*string {
	return s.QueryKeywords
}

func (s *RunSearchCaseFullTextResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCaseLevel(v string) *RunSearchCaseFullTextResponseBodyData {
	s.CaseLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCaseResult(v []*RunSearchCaseFullTextResponseBodyDataCaseResult) *RunSearchCaseFullTextResponseBodyData {
	s.CaseResult = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCurrentPage(v int32) *RunSearchCaseFullTextResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetPageSize(v int32) *RunSearchCaseFullTextResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetQuery(v string) *RunSearchCaseFullTextResponseBodyData {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetQueryKeywords(v []*string) *RunSearchCaseFullTextResponseBodyData {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetTotalCount(v int64) *RunSearchCaseFullTextResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) Validate() error {
	if s.CaseResult != nil {
		for _, item := range s.CaseResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchCaseFullTextResponseBodyDataCaseResult struct {
	CaseDomain *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain `json:"caseDomain,omitempty" xml:"caseDomain,omitempty" type:"Struct"`
	Mode       *string                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0.88
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResult) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) GetCaseDomain() *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	return s.CaseDomain
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) GetMode() *string {
	return s.Mode
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) GetSimilarity() *string {
	return s.Similarity
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetCaseDomain(v *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.CaseDomain = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetMode(v string) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.Mode = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetSimilarity(v string) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.Similarity = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) Validate() error {
	if s.CaseDomain != nil {
		if err := s.CaseDomain.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain struct {
	AbstractObj           *string                                                              `json:"abstractObj,omitempty" xml:"abstractObj,omitempty"`
	AppliedLaws           *string                                                              `json:"appliedLaws,omitempty" xml:"appliedLaws,omitempty"`
	BasicCase             *string                                                              `json:"basicCase,omitempty" xml:"basicCase,omitempty"`
	CaseBasic             *string                                                              `json:"caseBasic,omitempty" xml:"caseBasic,omitempty"`
	CaseCause             *string                                                              `json:"caseCause,omitempty" xml:"caseCause,omitempty"`
	CaseFeature           *string                                                              `json:"caseFeature,omitempty" xml:"caseFeature,omitempty"`
	CaseId                *string                                                              `json:"caseId,omitempty" xml:"caseId,omitempty"`
	CaseNo                *string                                                              `json:"caseNo,omitempty" xml:"caseNo,omitempty"`
	CaseSummary           *string                                                              `json:"caseSummary,omitempty" xml:"caseSummary,omitempty"`
	CaseTitle             *string                                                              `json:"caseTitle,omitempty" xml:"caseTitle,omitempty"`
	CaseType              *string                                                              `json:"caseType,omitempty" xml:"caseType,omitempty"`
	CloseCaseCause        *string                                                              `json:"closeCaseCause,omitempty" xml:"closeCaseCause,omitempty"`
	CourtFindOut          *string                                                              `json:"courtFindOut,omitempty" xml:"courtFindOut,omitempty"`
	CourtThink            *string                                                              `json:"courtThink,omitempty" xml:"courtThink,omitempty"`
	DataFrom              *string                                                              `json:"dataFrom,omitempty" xml:"dataFrom,omitempty"`
	DisputeFocus          *string                                                              `json:"disputeFocus,omitempty" xml:"disputeFocus,omitempty"`
	DisputeFocusTag       []*string                                                            `json:"disputeFocusTag,omitempty" xml:"disputeFocusTag,omitempty" type:"Repeated"`
	Disputedpoints        *string                                                              `json:"disputedpoints,omitempty" xml:"disputedpoints,omitempty"`
	DocumentType          *string                                                              `json:"documentType,omitempty" xml:"documentType,omitempty"`
	JudgReason            *string                                                              `json:"judgReason,omitempty" xml:"judgReason,omitempty"`
	Keyfacts              *string                                                              `json:"keyfacts,omitempty" xml:"keyfacts,omitempty"`
	LegalBasis            *string                                                              `json:"legalBasis,omitempty" xml:"legalBasis,omitempty"`
	Litigants             *string                                                              `json:"litigants,omitempty" xml:"litigants,omitempty"`
	LitigationParticipant *string                                                              `json:"litigationParticipant,omitempty" xml:"litigationParticipant,omitempty"`
	OpenCaseCause         *string                                                              `json:"openCaseCause,omitempty" xml:"openCaseCause,omitempty"`
	PreTrialProcess       *string                                                              `json:"preTrialProcess,omitempty" xml:"preTrialProcess,omitempty"`
	ReferLevel            *string                                                              `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	RefereeGist           *string                                                              `json:"refereeGist,omitempty" xml:"refereeGist,omitempty"`
	SourceContent         *string                                                              `json:"sourceContent,omitempty" xml:"sourceContent,omitempty"`
	TrialCourt            *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt `json:"trialCourt,omitempty" xml:"trialCourt,omitempty" type:"Struct"`
	// example:
	//
	// 2018-09-27
	TrialDate    *string `json:"trialDate,omitempty" xml:"trialDate,omitempty"`
	TrialLevel   *string `json:"trialLevel,omitempty" xml:"trialLevel,omitempty"`
	TrialProcess *string `json:"trialProcess,omitempty" xml:"trialProcess,omitempty"`
	TrialProgram *string `json:"trialProgram,omitempty" xml:"trialProgram,omitempty"`
	Verdict      *string `json:"verdict,omitempty" xml:"verdict,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetAbstractObj() *string {
	return s.AbstractObj
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetAppliedLaws() *string {
	return s.AppliedLaws
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetBasicCase() *string {
	return s.BasicCase
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseBasic() *string {
	return s.CaseBasic
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseCause() *string {
	return s.CaseCause
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseFeature() *string {
	return s.CaseFeature
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseId() *string {
	return s.CaseId
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseNo() *string {
	return s.CaseNo
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseSummary() *string {
	return s.CaseSummary
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseTitle() *string {
	return s.CaseTitle
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCaseType() *string {
	return s.CaseType
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCloseCaseCause() *string {
	return s.CloseCaseCause
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCourtFindOut() *string {
	return s.CourtFindOut
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetCourtThink() *string {
	return s.CourtThink
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetDataFrom() *string {
	return s.DataFrom
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetDisputeFocus() *string {
	return s.DisputeFocus
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetDisputeFocusTag() []*string {
	return s.DisputeFocusTag
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetDisputedpoints() *string {
	return s.Disputedpoints
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetDocumentType() *string {
	return s.DocumentType
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetJudgReason() *string {
	return s.JudgReason
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetKeyfacts() *string {
	return s.Keyfacts
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetLegalBasis() *string {
	return s.LegalBasis
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetLitigants() *string {
	return s.Litigants
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetLitigationParticipant() *string {
	return s.LitigationParticipant
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetOpenCaseCause() *string {
	return s.OpenCaseCause
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetPreTrialProcess() *string {
	return s.PreTrialProcess
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetReferLevel() *string {
	return s.ReferLevel
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetRefereeGist() *string {
	return s.RefereeGist
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetSourceContent() *string {
	return s.SourceContent
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetTrialCourt() *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	return s.TrialCourt
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetTrialDate() *string {
	return s.TrialDate
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetTrialLevel() *string {
	return s.TrialLevel
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetTrialProcess() *string {
	return s.TrialProcess
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetTrialProgram() *string {
	return s.TrialProgram
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GetVerdict() *string {
	return s.Verdict
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetAbstractObj(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.AbstractObj = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetAppliedLaws(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.AppliedLaws = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetBasicCase(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.BasicCase = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseBasic(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseBasic = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseFeature(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseFeature = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseId(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseId = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseNo(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseNo = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseSummary(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseSummary = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseTitle(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseTitle = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseType(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseType = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCloseCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CloseCaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCourtFindOut(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CourtFindOut = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCourtThink(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CourtThink = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDataFrom(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DataFrom = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputeFocus(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DisputeFocus = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputeFocusTag(v []*string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DisputeFocusTag = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputedpoints(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Disputedpoints = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDocumentType(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DocumentType = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetJudgReason(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.JudgReason = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetKeyfacts(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Keyfacts = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLegalBasis(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.LegalBasis = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLitigants(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Litigants = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLitigationParticipant(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.LitigationParticipant = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetOpenCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.OpenCaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetPreTrialProcess(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.PreTrialProcess = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetReferLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetRefereeGist(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.RefereeGist = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetSourceContent(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.SourceContent = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialCourt(v *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialCourt = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialDate(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialDate = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialProcess(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialProcess = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialProgram(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialProgram = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetVerdict(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Verdict = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) Validate() error {
	if s.TrialCourt != nil {
		if err := s.TrialCourt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt struct {
	City        *string `json:"city,omitempty" xml:"city,omitempty"`
	CommonLevel *string `json:"commonLevel,omitempty" xml:"commonLevel,omitempty"`
	Country     *string `json:"country,omitempty" xml:"country,omitempty"`
	County      *string `json:"county,omitempty" xml:"county,omitempty"`
	District    *string `json:"district,omitempty" xml:"district,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Province    *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// “”
	SpecialLevel *string `json:"specialLevel,omitempty" xml:"specialLevel,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetCity() *string {
	return s.City
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetCommonLevel() *string {
	return s.CommonLevel
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetCountry() *string {
	return s.Country
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetCounty() *string {
	return s.County
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetDistrict() *string {
	return s.District
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetName() *string {
	return s.Name
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetProvince() *string {
	return s.Province
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GetSpecialLevel() *string {
	return s.SpecialLevel
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCity(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.City = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCommonLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.CommonLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCountry(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Country = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCounty(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.County = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetDistrict(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.District = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetName(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Name = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetProvince(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Province = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetSpecialLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.SpecialLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) Validate() error {
	return dara.Validate(s)
}
