// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliverySLRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *CreateLogDeliverySLRRequest
	GetConsoleSessionId() *string
}

type CreateLogDeliverySLRRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s CreateLogDeliverySLRRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliverySLRRequest) GoString() string {
	return s.String()
}

func (s *CreateLogDeliverySLRRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CreateLogDeliverySLRRequest) SetConsoleSessionId(v string) *CreateLogDeliverySLRRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CreateLogDeliverySLRRequest) Validate() error {
	return dara.Validate(s)
}
