/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTemplateData Request to create a Template.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.CreateTemplateRequest
type VmwareTanzuManageV1alpha1PolicyTemplateData struct {

	// Template to create.
	Template *VmwareTanzuManageV1alpha1PolicyTemplate `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplateData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1PolicyTemplateListData Request to create a Template.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.ListTemplatesResponse
type VmwareTanzuManageV1alpha1PolicyTemplateListData struct {

	// List of templates.
	Templates []*VmwareTanzuManageV1alpha1PolicyTemplate `json:"templates"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateListData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateListData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplateListData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// !!! NOT GENERATED BY SWAGGER !!!.

type ListCustomTemplatesRequest struct {
	// Scope can be provider or cluster.
	TemplateName string `json:"templateName"`

	// Sort results by.
	SortBy string `json:"sortBy,omitempty"`

	// Query to run against the API.
	Query string `json:"query,omitempty"`

	// Include Total.
	IncludeTotalCount bool `json:"includeTotal"`
}

// MarshalBinary interface implementation.
func (m *ListCustomTemplatesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *ListCustomTemplatesRequest) UnmarshalBinary(b []byte) error {
	var res ListCustomTemplatesRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
