// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSourceLocationResponseBody
	GetRequestId() *string
	SetSourceLocation(v *ChannelAssemblySourceLocation) *UpdateSourceLocationResponseBody
	GetSourceLocation() *ChannelAssemblySourceLocation
}

type UpdateSourceLocationResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source location information.
	SourceLocation *ChannelAssemblySourceLocation `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s UpdateSourceLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSourceLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSourceLocationResponseBody) GetSourceLocation() *ChannelAssemblySourceLocation {
	return s.SourceLocation
}

func (s *UpdateSourceLocationResponseBody) SetRequestId(v string) *UpdateSourceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSourceLocationResponseBody) SetSourceLocation(v *ChannelAssemblySourceLocation) *UpdateSourceLocationResponseBody {
	s.SourceLocation = v
	return s
}

func (s *UpdateSourceLocationResponseBody) Validate() error {
	if s.SourceLocation != nil {
		if err := s.SourceLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
