// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstancePreivewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *InstancePreivewRequest
	GetConsoleSessionId() *string
	SetTags(v string) *InstancePreivewRequest
	GetTags() *string
}

type InstancePreivewRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s InstancePreivewRequest) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewRequest) GoString() string {
	return s.String()
}

func (s *InstancePreivewRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *InstancePreivewRequest) GetTags() *string {
	return s.Tags
}

func (s *InstancePreivewRequest) SetConsoleSessionId(v string) *InstancePreivewRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *InstancePreivewRequest) SetTags(v string) *InstancePreivewRequest {
	s.Tags = &v
	return s
}

func (s *InstancePreivewRequest) Validate() error {
	return dara.Validate(s)
}
