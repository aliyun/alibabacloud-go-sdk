// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityDefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaEntityDef(v *MetaEntityDef) *GetMetaEntityDefResponseBody
	GetMetaEntityDef() *MetaEntityDef
	SetRequestId(v string) *GetMetaEntityDefResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaEntityDefResponseBody
	GetSuccess() *bool
}

type GetMetaEntityDefResponseBody struct {
	MetaEntityDef *MetaEntityDef `json:"MetaEntityDef,omitempty" xml:"MetaEntityDef,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C636A747-7E4E-594D-94CD-2B4F8A9A9A63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaEntityDefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityDefResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaEntityDefResponseBody) GetMetaEntityDef() *MetaEntityDef {
	return s.MetaEntityDef
}

func (s *GetMetaEntityDefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaEntityDefResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaEntityDefResponseBody) SetMetaEntityDef(v *MetaEntityDef) *GetMetaEntityDefResponseBody {
	s.MetaEntityDef = v
	return s
}

func (s *GetMetaEntityDefResponseBody) SetRequestId(v string) *GetMetaEntityDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaEntityDefResponseBody) SetSuccess(v bool) *GetMetaEntityDefResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaEntityDefResponseBody) Validate() error {
	if s.MetaEntityDef != nil {
		if err := s.MetaEntityDef.Validate(); err != nil {
			return err
		}
	}
	return nil
}
