// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSourceResponseBody
	GetRequestId() *string
	SetSource(v *ChannelAssemblySource) *UpdateSourceResponseBody
	GetSource() *ChannelAssemblySource
}

type UpdateSourceResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source information.
	Source *ChannelAssemblySource `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSourceResponseBody) GetSource() *ChannelAssemblySource {
	return s.Source
}

func (s *UpdateSourceResponseBody) SetRequestId(v string) *UpdateSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSourceResponseBody) SetSource(v *ChannelAssemblySource) *UpdateSourceResponseBody {
	s.Source = v
	return s
}

func (s *UpdateSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
