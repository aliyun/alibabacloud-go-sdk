// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProjectsResponseBody
	GetRequestId() *string
	SetResults(v []*DeleteProjectsResponseBodyResults) *DeleteProjectsResponseBody
	GetResults() []*DeleteProjectsResponseBodyResults
}

type DeleteProjectsResponseBody struct {
	// The unique identifier that Alibaba Cloud generated for the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The call results of the operation.
	Results []*DeleteProjectsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DeleteProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectsResponseBody) GetResults() []*DeleteProjectsResponseBodyResults {
	return s.Results
}

func (s *DeleteProjectsResponseBody) SetRequestId(v string) *DeleteProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectsResponseBody) SetResults(v []*DeleteProjectsResponseBodyResults) *DeleteProjectsResponseBody {
	s.Results = v
	return s
}

func (s *DeleteProjectsResponseBody) Validate() error {
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

type DeleteProjectsResponseBodyResults struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 2
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DeleteProjectsResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DeleteProjectsResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DeleteProjectsResponseBodyResults) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteProjectsResponseBodyResults) SetCode(v string) *DeleteProjectsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DeleteProjectsResponseBodyResults) SetMessage(v string) *DeleteProjectsResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DeleteProjectsResponseBodyResults) SetProjectId(v string) *DeleteProjectsResponseBodyResults {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
