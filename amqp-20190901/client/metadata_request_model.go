// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *MetadataRequest
	GetConsoleSessionId() *string
}

type MetadataRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s MetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s MetadataRequest) GoString() string {
	return s.String()
}

func (s *MetadataRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *MetadataRequest) SetConsoleSessionId(v string) *MetadataRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *MetadataRequest) Validate() error {
	return dara.Validate(s)
}
