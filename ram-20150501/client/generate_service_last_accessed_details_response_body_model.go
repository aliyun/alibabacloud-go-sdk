// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServiceLastAccessedDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GenerateServiceLastAccessedDetailsResponseBody
	GetJobId() *string
	SetRequestId(v string) *GenerateServiceLastAccessedDetailsResponseBody
	GetRequestId() *string
}

type GenerateServiceLastAccessedDetailsResponseBody struct {
	// The ID of the permission audit task. You can use this ID to call the GetServiceLastAccessedDetails operation to query the report generation status.
	//
	// example:
	//
	// ba536aab-84c6-43af-8bcb-aa2cae2a34e4
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEE5D280-31A0-1638-91CD-37EBDD282E7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateServiceLastAccessedDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateServiceLastAccessedDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateServiceLastAccessedDetailsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GenerateServiceLastAccessedDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateServiceLastAccessedDetailsResponseBody) SetJobId(v string) *GenerateServiceLastAccessedDetailsResponseBody {
	s.JobId = &v
	return s
}

func (s *GenerateServiceLastAccessedDetailsResponseBody) SetRequestId(v string) *GenerateServiceLastAccessedDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateServiceLastAccessedDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
