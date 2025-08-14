// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSourceLocationResponseBody
	GetRequestId() *string
	SetSourceLocation(v *ChannelAssemblySourceLocation) *CreateSourceLocationResponseBody
	GetSourceLocation() *ChannelAssemblySourceLocation
}

type CreateSourceLocationResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source location information.
	SourceLocation *ChannelAssemblySourceLocation `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s CreateSourceLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSourceLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSourceLocationResponseBody) GetSourceLocation() *ChannelAssemblySourceLocation {
	return s.SourceLocation
}

func (s *CreateSourceLocationResponseBody) SetRequestId(v string) *CreateSourceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSourceLocationResponseBody) SetSourceLocation(v *ChannelAssemblySourceLocation) *CreateSourceLocationResponseBody {
	s.SourceLocation = v
	return s
}

func (s *CreateSourceLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
