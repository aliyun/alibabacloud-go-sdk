// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPolarClawSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SearchPolarClawSkillsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *SearchPolarClawSkillsResponseBody
	GetCode() *int32
	SetMessage(v string) *SearchPolarClawSkillsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchPolarClawSkillsResponseBody
	GetRequestId() *string
	SetResults(v []*SearchPolarClawSkillsResponseBodyResults) *SearchPolarClawSkillsResponseBody
	GetResults() []*SearchPolarClawSkillsResponseBodyResults
}

type SearchPolarClawSkillsResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of search results.
	Results []*SearchPolarClawSkillsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s SearchPolarClawSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchPolarClawSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPolarClawSkillsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SearchPolarClawSkillsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchPolarClawSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchPolarClawSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchPolarClawSkillsResponseBody) GetResults() []*SearchPolarClawSkillsResponseBodyResults {
	return s.Results
}

func (s *SearchPolarClawSkillsResponseBody) SetApplicationId(v string) *SearchPolarClawSkillsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBody) SetCode(v int32) *SearchPolarClawSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBody) SetMessage(v string) *SearchPolarClawSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBody) SetRequestId(v string) *SearchPolarClawSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBody) SetResults(v []*SearchPolarClawSkillsResponseBodyResults) *SearchPolarClawSkillsResponseBody {
	s.Results = v
	return s
}

func (s *SearchPolarClawSkillsResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchPolarClawSkillsResponseBodyResults struct {
	// The display name.
	//
	// example:
	//
	// RDS Copilot
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The relevance score.
	//
	// example:
	//
	// 0.95
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The skill identifier.
	//
	// example:
	//
	// alibacloud-rds-copilot
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The brief description.
	//
	// example:
	//
	// A copilot for RDS
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The UNIX timestamp of the last update, in seconds.
	//
	// example:
	//
	// 1716000000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 1.2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchPolarClawSkillsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SearchPolarClawSkillsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetDisplayName() *string {
	return s.DisplayName
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetScore() *float64 {
	return s.Score
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetSlug() *string {
	return s.Slug
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetSummary() *string {
	return s.Summary
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *SearchPolarClawSkillsResponseBodyResults) GetVersion() *string {
	return s.Version
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetDisplayName(v string) *SearchPolarClawSkillsResponseBodyResults {
	s.DisplayName = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetScore(v float64) *SearchPolarClawSkillsResponseBodyResults {
	s.Score = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetSlug(v string) *SearchPolarClawSkillsResponseBodyResults {
	s.Slug = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetSummary(v string) *SearchPolarClawSkillsResponseBodyResults {
	s.Summary = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetUpdatedAt(v int64) *SearchPolarClawSkillsResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) SetVersion(v string) *SearchPolarClawSkillsResponseBodyResults {
	s.Version = &v
	return s
}

func (s *SearchPolarClawSkillsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
