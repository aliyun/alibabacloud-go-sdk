// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserCustomLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigIds(v *ListUserCustomLogConfigResponseBodyConfigIds) *ListUserCustomLogConfigResponseBody
	GetConfigIds() *ListUserCustomLogConfigResponseBodyConfigIds
	SetRequestId(v string) *ListUserCustomLogConfigResponseBody
	GetRequestId() *string
}

type ListUserCustomLogConfigResponseBody struct {
	// The list of log configuration IDs.
	ConfigIds *ListUserCustomLogConfigResponseBodyConfigIds `json:"ConfigIds,omitempty" xml:"ConfigIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 95D5B69F-8AEC-419B-8F3A-612B35032B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserCustomLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponseBody) GetConfigIds() *ListUserCustomLogConfigResponseBodyConfigIds {
	return s.ConfigIds
}

func (s *ListUserCustomLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserCustomLogConfigResponseBody) SetConfigIds(v *ListUserCustomLogConfigResponseBodyConfigIds) *ListUserCustomLogConfigResponseBody {
	s.ConfigIds = v
	return s
}

func (s *ListUserCustomLogConfigResponseBody) SetRequestId(v string) *ListUserCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserCustomLogConfigResponseBody) Validate() error {
	if s.ConfigIds != nil {
		if err := s.ConfigIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserCustomLogConfigResponseBodyConfigIds struct {
	ConfigId []*string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty" type:"Repeated"`
}

func (s ListUserCustomLogConfigResponseBodyConfigIds) String() string {
	return dara.Prettify(s)
}

func (s ListUserCustomLogConfigResponseBodyConfigIds) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponseBodyConfigIds) GetConfigId() []*string {
	return s.ConfigId
}

func (s *ListUserCustomLogConfigResponseBodyConfigIds) SetConfigId(v []*string) *ListUserCustomLogConfigResponseBodyConfigIds {
	s.ConfigId = v
	return s
}

func (s *ListUserCustomLogConfigResponseBodyConfigIds) Validate() error {
	return dara.Validate(s)
}
