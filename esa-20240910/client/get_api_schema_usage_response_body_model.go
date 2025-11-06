// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiSchemaUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetApiSchemaUsageResponseBody
	GetInstanceId() *string
	SetInstanceUsage(v int32) *GetApiSchemaUsageResponseBody
	GetInstanceUsage() *int32
	SetRequestId(v string) *GetApiSchemaUsageResponseBody
	GetRequestId() *string
	SetUsages(v []*GetApiSchemaUsageResponseBodyUsages) *GetApiSchemaUsageResponseBody
	GetUsages() []*GetApiSchemaUsageResponseBodyUsages
}

type GetApiSchemaUsageResponseBody struct {
	// example:
	//
	// esa-site-agknce3n****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	InstanceUsage *int32 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 952ea16b-1f05-4a76-bb32-420282d8****
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Usages    []*GetApiSchemaUsageResponseBodyUsages `json:"Usages,omitempty" xml:"Usages,omitempty" type:"Repeated"`
}

func (s GetApiSchemaUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiSchemaUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiSchemaUsageResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApiSchemaUsageResponseBody) GetInstanceUsage() *int32 {
	return s.InstanceUsage
}

func (s *GetApiSchemaUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiSchemaUsageResponseBody) GetUsages() []*GetApiSchemaUsageResponseBodyUsages {
	return s.Usages
}

func (s *GetApiSchemaUsageResponseBody) SetInstanceId(v string) *GetApiSchemaUsageResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetApiSchemaUsageResponseBody) SetInstanceUsage(v int32) *GetApiSchemaUsageResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *GetApiSchemaUsageResponseBody) SetRequestId(v string) *GetApiSchemaUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiSchemaUsageResponseBody) SetUsages(v []*GetApiSchemaUsageResponseBodyUsages) *GetApiSchemaUsageResponseBody {
	s.Usages = v
	return s
}

func (s *GetApiSchemaUsageResponseBody) Validate() error {
	if s.Usages != nil {
		for _, item := range s.Usages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiSchemaUsageResponseBodyUsages struct {
	// example:
	//
	// 40000449
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetApiSchemaUsageResponseBodyUsages) String() string {
	return dara.Prettify(s)
}

func (s GetApiSchemaUsageResponseBodyUsages) GoString() string {
	return s.String()
}

func (s *GetApiSchemaUsageResponseBodyUsages) GetId() *int64 {
	return s.Id
}

func (s *GetApiSchemaUsageResponseBodyUsages) GetName() *string {
	return s.Name
}

func (s *GetApiSchemaUsageResponseBodyUsages) GetUsage() *int32 {
	return s.Usage
}

func (s *GetApiSchemaUsageResponseBodyUsages) SetId(v int64) *GetApiSchemaUsageResponseBodyUsages {
	s.Id = &v
	return s
}

func (s *GetApiSchemaUsageResponseBodyUsages) SetName(v string) *GetApiSchemaUsageResponseBodyUsages {
	s.Name = &v
	return s
}

func (s *GetApiSchemaUsageResponseBodyUsages) SetUsage(v int32) *GetApiSchemaUsageResponseBodyUsages {
	s.Usage = &v
	return s
}

func (s *GetApiSchemaUsageResponseBodyUsages) Validate() error {
	return dara.Validate(s)
}
