// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkInfoByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AccountLinkInfo) *GetLinkInfoByUserIdResponseBody
	GetItems() []*AccountLinkInfo
}

type GetLinkInfoByUserIdResponseBody struct {
	// The information about the users.
	Items []*AccountLinkInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s GetLinkInfoByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLinkInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdResponseBody) GetItems() []*AccountLinkInfo {
	return s.Items
}

func (s *GetLinkInfoByUserIdResponseBody) SetItems(v []*AccountLinkInfo) *GetLinkInfoByUserIdResponseBody {
	s.Items = v
	return s
}

func (s *GetLinkInfoByUserIdResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
