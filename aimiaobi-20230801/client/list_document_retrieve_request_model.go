// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentRetrieveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *ListDocumentRetrieveRequest
	GetContentType() *string
	SetElementScope(v string) *ListDocumentRetrieveRequest
	GetElementScope() *string
	SetEndDate(v string) *ListDocumentRetrieveRequest
	GetEndDate() *string
	SetMaxResults(v int32) *ListDocumentRetrieveRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocumentRetrieveRequest
	GetNextToken() *string
	SetOffice(v string) *ListDocumentRetrieveRequest
	GetOffice() *string
	SetQuery(v string) *ListDocumentRetrieveRequest
	GetQuery() *string
	SetRegion(v string) *ListDocumentRetrieveRequest
	GetRegion() *string
	SetSource(v string) *ListDocumentRetrieveRequest
	GetSource() *string
	SetStartDate(v string) *ListDocumentRetrieveRequest
	GetStartDate() *string
	SetSubContentType(v string) *ListDocumentRetrieveRequest
	GetSubContentType() *string
	SetSubjectClassify(v string) *ListDocumentRetrieveRequest
	GetSubjectClassify() *string
	SetWordSize(v string) *ListDocumentRetrieveRequest
	GetWordSize() *string
	SetWorkspaceId(v string) *ListDocumentRetrieveRequest
	GetWorkspaceId() *string
}

type ListDocumentRetrieveRequest struct {
	// Document type. Valid values: 0 (default): All types. 1: Government documents. 2: Important articles. 5: Policy interpretation. 6: Legal provisions. 7: Regulations and rules. 8: General Secretary.
	//
	// example:
	//
	// 1
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Search scope. Valid values: 1: Title only. 0: Full text (title and content). Default is 0.
	//
	// example:
	//
	// 0
	ElementScope *string `json:"ElementScope,omitempty" xml:"ElementScope,omitempty"`
	// End date of issuance in yyyy-MM-dd format.
	//
	// example:
	//
	// 2025-07-03
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Maximum number of results to return.
	//
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next page of results.
	//
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Issuing agency.
	//
	// example:
	//
	// 国务院办公室
	Office *string `json:"Office,omitempty" xml:"Office,omitempty"`
	// Search condition.
	//
	// example:
	//
	// 检索Query
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region. Enter a province or city, such as Jilin Province or Beijing Municipality.
	//
	// example:
	//
	// 北京市
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Source. Valid values: 0: Internal (within your organization). 1: External (outside your organization).
	//
	// example:
	//
	// 1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Start date of issuance in yyyy-MM-dd format.
	//
	// example:
	//
	// 2025-10-10
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// - Secondary classification of document type.
	//
	//   - When the document type is an official document: -1: Other; 0: Resolution; 1: Decision; 2: Order; 3: Bulletin; 4: Public Notice; 5: Notice; 6: Opinion; 7: Notification; 8: Circular; 9: Report; 10: Request for Instructions; 11: Approval; 12: Motion; 13: Letter
	//
	//   - 14: Summary
	//
	//   - When the article type is important articles: 1: important commentary 2: important theory 3: other articles
	//
	//   - When the document genre is rules and regulations: 3: Administrative regulations 4: Supervisory regulations 5: Local regulations 7: Departmental rules 8: Others 9: Party constitution and regulations
	//
	//   - When the article genre is a legal provision: 1: Constitution 2: Law 6: Judicial Interpretation
	//
	// example:
	//
	// 1
	SubContentType *string `json:"SubContentType,omitempty" xml:"SubContentType,omitempty"`
	// Supported classifications:
	//
	// | Level 1 category                                                        | Level 2 category                                                                                                                                                                                                                                  |
	//
	// | ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//
	// | National defense and international cooperation                          | National defense. Foreign affairs. Military affairs. Work related to Hong Kong, Macao, Taiwan, and overseas Chinese.                                                                                                                              |
	//
	// | Comprehensive administration                                            | The 20th National Congress of the Communist Party of China. Government transparency and supervision. Joint administration. Party building. Conferences and proposals. Government document management. Other administrative matters.               |
	//
	// | State Council organizational structure                                  | State Council. General Office of the State Council. State Council agencies.                                                                                                                                                                       |
	//
	// | Administrative and market regulation                                    | Administrative regulation. Credit regulation. Product quality supervision. Work safety supervision. Market regulation.                                                                                                                            |
	//
	// | Economic management                                                     | National economy. Market economy. Economic system reform. State-owned asset supervision.                                                                                                                                                          |
	//
	// | Finance, banking, commerce, and customs                                 | Finance. Banking. Auditing. Commerce. Customs.                                                                                                                                                                                                    |
	//
	// | Personnel and social security                                           | Personnel work. Population and family planning. Work related to women and children. Poverty alleviation. Disaster reduction and relief. Public services. Social welfare and assistance. Preferential treatment and resettlement. Social security. |
	//
	// | Public security and social management                                   | Public security. Safety. Justice. Fire control. Ethnic affairs. Religion.                                                                                                                                                                         |
	//
	// | Science, education, culture, and sports                                 | Culture. Scientific and technological innovation. Education. Intellectual property. Press and publishing. Radio, television, and the Internet. Sports. Tourism.                                                                                   |
	//
	// | Healthcare                                                              | Health. Medical care. Veterinary medicine.                                                                                                                                                                                                        |
	//
	// | Urban-rural development and industrial growth                           | Urban and rural development. Industry. Transportation.                                                                                                                                                                                            |
	//
	// | Natural resources and environmental protection                          | Land and energy resources. Civil engineering. Meteorology. Environmental protection.                                                                                                                                                              |
	//
	// | Agriculture, forestry, water resources, fisheries, and animal husbandry | Agriculture. Forestry. Water resources. Fisheries. Animal husbandry.                                                                                                                                                                              |
	//
	// | Others                                                                  | Others.                                                                                                                                                                                                                                           |
	//
	// example:
	//
	// 国防和交流合作事务
	SubjectClassify *string `json:"SubjectClassify,omitempty" xml:"SubjectClassify,omitempty"`
	// Document number.
	//
	// example:
	//
	// 宁民规〔2020〕5号
	WordSize *string `json:"WordSize,omitempty" xml:"WordSize,omitempty"`
	// Unique identifier of the Model Studio workspace. For more information, see [Get workspaceId](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDocumentRetrieveRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentRetrieveRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentRetrieveRequest) GetContentType() *string {
	return s.ContentType
}

func (s *ListDocumentRetrieveRequest) GetElementScope() *string {
	return s.ElementScope
}

func (s *ListDocumentRetrieveRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListDocumentRetrieveRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentRetrieveRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentRetrieveRequest) GetOffice() *string {
	return s.Office
}

func (s *ListDocumentRetrieveRequest) GetQuery() *string {
	return s.Query
}

func (s *ListDocumentRetrieveRequest) GetRegion() *string {
	return s.Region
}

func (s *ListDocumentRetrieveRequest) GetSource() *string {
	return s.Source
}

func (s *ListDocumentRetrieveRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListDocumentRetrieveRequest) GetSubContentType() *string {
	return s.SubContentType
}

func (s *ListDocumentRetrieveRequest) GetSubjectClassify() *string {
	return s.SubjectClassify
}

func (s *ListDocumentRetrieveRequest) GetWordSize() *string {
	return s.WordSize
}

func (s *ListDocumentRetrieveRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDocumentRetrieveRequest) SetContentType(v string) *ListDocumentRetrieveRequest {
	s.ContentType = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetElementScope(v string) *ListDocumentRetrieveRequest {
	s.ElementScope = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetEndDate(v string) *ListDocumentRetrieveRequest {
	s.EndDate = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetMaxResults(v int32) *ListDocumentRetrieveRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetNextToken(v string) *ListDocumentRetrieveRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetOffice(v string) *ListDocumentRetrieveRequest {
	s.Office = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetQuery(v string) *ListDocumentRetrieveRequest {
	s.Query = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetRegion(v string) *ListDocumentRetrieveRequest {
	s.Region = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetSource(v string) *ListDocumentRetrieveRequest {
	s.Source = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetStartDate(v string) *ListDocumentRetrieveRequest {
	s.StartDate = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetSubContentType(v string) *ListDocumentRetrieveRequest {
	s.SubContentType = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetSubjectClassify(v string) *ListDocumentRetrieveRequest {
	s.SubjectClassify = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetWordSize(v string) *ListDocumentRetrieveRequest {
	s.WordSize = &v
	return s
}

func (s *ListDocumentRetrieveRequest) SetWorkspaceId(v string) *ListDocumentRetrieveRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDocumentRetrieveRequest) Validate() error {
	return dara.Validate(s)
}
