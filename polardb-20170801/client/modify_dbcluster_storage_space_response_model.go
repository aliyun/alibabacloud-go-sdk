// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStorageSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterStorageSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterStorageSpaceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterStorageSpaceResponseBody) *ModifyDBClusterStorageSpaceResponse
	GetBody() *ModifyDBClusterStorageSpaceResponseBody
}

type ModifyDBClusterStorageSpaceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterStorageSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterStorageSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStorageSpaceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStorageSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterStorageSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterStorageSpaceResponse) GetBody() *ModifyDBClusterStorageSpaceResponseBody {
	return s.Body
}

func (s *ModifyDBClusterStorageSpaceResponse) SetHeaders(v map[string]*string) *ModifyDBClusterStorageSpaceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponse) SetStatusCode(v int32) *ModifyDBClusterStorageSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponse) SetBody(v *ModifyDBClusterStorageSpaceResponseBody) *ModifyDBClusterStorageSpaceResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponse) Validate() error {
	return dara.Validate(s)
}
